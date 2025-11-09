import React, { useState } from 'react';
import { Form, Input, InputNumber, DatePicker, Button, Card, Modal, message, Space, Typography, Row, Col } from 'antd';
import { GiftOutlined, QrcodeOutlined, LinkOutlined } from '@ant-design/icons';
import dayjs from 'dayjs';
import api from '../services/api';
import SpinWheel from './SpinWheel';  // ⭐ 引入輪盤組件

const { Title, Text, Paragraph } = Typography;

function IssuePanel() {
  const [form] = Form.useForm();
  const [loading, setLoading] = useState(false);
  const [qrData, setQrData] = useState(null);
  const [modalVisible, setModalVisible] = useState(false);

  // ⭐ 處理輪盤結果
  const handleSpinResult = (amount) => {
    form.setFieldsValue({ discount_amount: amount });
    message.success(`🎉 幸運輪盤結果: ${amount} 元!`);
  };

  const onFinish = async (values) => {
    setLoading(true);
    try {
      const data = {
        customer_name: values.customer_name,
        discount_amount: values.discount_amount,
        expired_date: values.expired_date.format('YYYY-MM-DD'),
      };

      const response = await api.issueCoupon(data);
      
      setQrData(response.data);
      setModalVisible(true);
      message.success(response.data.message);
      form.resetFields();
      
    } catch (error) {
      message.error('發行失敗: ' + (error.response?.data?.error || error.message));
    } finally {
      setLoading(false);
    }
  };

  const handleCopyDeepLink = () => {
    if (qrData?.deep_link) {
      navigator.clipboard.writeText(qrData.deep_link);
      message.success('Deep Link 已複製到剪貼簿');
    }
  };

  return (
    <div className="issue-panel">
      <Card className="issue-card">
        <Title level={3}>
          <GiftOutlined /> 發行折價券
        </Title>
        
        <Form
          form={form}
          layout="vertical"
          onFinish={onFinish}
          initialValues={{
            discount_amount: 200,
            expired_date: dayjs().add(7, 'days'),
          }}
        >
          <Form.Item
            label="消費者姓名"
            name="customer_name"
            rules={[
              { required: true, message: '請輸入姓名' },
              { pattern: /^[\u4e00-\u9fa5]+$/, message: '只能輸入中文' }
            ]}
          >
            <Input 
              placeholder="請輸入消費者姓名 (僅限中文)" 
              size="large"
            />
          </Form.Item>

          {/* ⭐ 修改:折扣金額欄位 + 輪盤按鈕 */}
          <Form.Item
            label={
              <Space>
                <span>折扣金額 (新台幣)</span>
                <Text type="secondary" style={{ fontSize: '12px' }}>
                  (可使用幸運輪盤隨機決定)
                </Text>
              </Space>
            }
          >
            <Row gutter={16}>
              <Col span={16}>
                <Form.Item
                  name="discount_amount"
                  noStyle
                  rules={[
                    { required: true, message: '請輸入折扣金額' },
                    { type: 'number', min: 100, max: 999, message: '金額必須在 100-999 之間' }
                  ]}
                >
                  <InputNumber
                    style={{ width: '100%' }}
                    size="large"
                    min={100}
                    max={999}
                    formatter={value => `$ ${value}`.replace(/\B(?=(\d{3})+(?!\d))/g, ',')}
                    parser={value => value.replace(/\$\s?|(,*)/g, '')}
                  />
                </Form.Item>
              </Col>
              <Col span={8}>
                {/* ⭐ 幸運輪盤按鈕 */}
                <SpinWheel onResult={handleSpinResult} />
              </Col>
            </Row>
          </Form.Item>

          <Form.Item
            label="到期日期"
            name="expired_date"
            rules={[{ required: true, message: '請選擇到期日期' }]}
          >
            <DatePicker 
              style={{ width: '100%' }}
              size="large"
              format="YYYY-MM-DD"
              disabledDate={(current) => current && current < dayjs().startOf('day')}
            />
          </Form.Item>

          <Form.Item>
            <Button 
              type="primary" 
              htmlType="submit" 
              size="large"
              loading={loading}
              block
              icon={<QrcodeOutlined />}
            >
              產生 QR Code
            </Button>
          </Form.Item>
        </Form>
      </Card>

      <Modal
        title={<><QrcodeOutlined /> 折價券 QR Code</>}
        open={modalVisible}
        onCancel={() => setModalVisible(false)}
        footer={[
          <Button key="copy" icon={<LinkOutlined />} onClick={handleCopyDeepLink}>
            複製 Deep Link
          </Button>,
          <Button key="close" type="primary" onClick={() => setModalVisible(false)}>
            關閉
          </Button>
        ]}
        width={500}
      >
        {qrData && (
          <div style={{ textAlign: 'center' }}>
            <img 
              src={qrData.qr_code} 
              alt="Coupon QR Code"
              style={{ 
                width: '300px', 
                height: '300px',
                margin: '20px auto',
                display: 'block',
                border: '1px solid #d9d9d9',
                borderRadius: '4px'
              }}
            />
            <Paragraph>
              <Text strong>Transaction ID:</Text><br />
              <Text copyable code>{qrData.transaction_id}</Text>
            </Paragraph>
            <Paragraph type="secondary">
              請使用數位憑證皮夾 APP 掃描此 QR Code 以下載折價券
            </Paragraph>
          </div>
        )}
      </Modal>
    </div>
  );
}

export default IssuePanel;