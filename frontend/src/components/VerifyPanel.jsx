import React, { useState } from 'react';
import { Form, InputNumber, Button, Card, Statistic, Row, Col, Alert, Modal, message, Space, Typography, Divider, Steps } from 'antd';
import { CheckCircleOutlined, DollarOutlined, TagOutlined, QrcodeOutlined, ScanOutlined, LinkOutlined } from '@ant-design/icons';
import api from '../services/api';

const { Title, Text, Paragraph } = Typography;
const { Step } = Steps;

function VerifyPanel() {
  const [form] = Form.useForm();
  const [loading, setLoading] = useState(false);
  const [currentStep, setCurrentStep] = useState(0);
  
  // 驗證 QR Code 資料
  const [verifyQRData, setVerifyQRData] = useState(null);
  const [qrModalVisible, setQrModalVisible] = useState(false);
  
  // 驗證結果
  const [verifyResult, setVerifyResult] = useState(null);

  // 步驟 1: 產生驗證 QR Code
  const handleGenerateQR = async () => {
    setLoading(true);
    try {
      const response = await api.generateVerifyQR();
      setVerifyQRData(response.data.data);
      setQrModalVisible(true);
      setCurrentStep(1);
      message.success('驗證 QR Code 已產生!');
    } catch (error) {
      message.error('產生 QR Code 失敗: ' + (error.response?.data?.error || error.message));
    } finally {
      setLoading(false);
    }
  };

  // 步驟 2: 計算折扣(消費者掃描 QR Code 後)
  const onFinish = async (values) => {
    if (!verifyQRData) {
      message.error('請先產生驗證 QR Code!');
      return;
    }

    setLoading(true);
    try {
      const data = {
        transaction_id: verifyQRData.transactionId,
        original_amount: values.original_amount,
      };

      const response = await api.calculateDiscount(data);
      setVerifyResult(response.data);
      setCurrentStep(2);
      message.success('驗證成功!');
      
    } catch (error) {
      const errorMsg = error.response?.data?.error || error.message;
      if (errorMsg.includes('尚未上傳') || errorMsg.includes('not found')) {
        message.warning('消費者尚未掃描 QR Code,請稍後再試');
      } else {
        message.error('驗證失敗: ' + errorMsg);
      }
      setVerifyResult(null);
    } finally {
      setLoading(false);
    }
  };

  const resetForm = () => {
    form.resetFields();
    setVerifyResult(null);
    setVerifyQRData(null);
    setCurrentStep(0);
  };

  const handleCopyTransactionId = () => {
    if (verifyQRData?.transactionId) {
      navigator.clipboard.writeText(verifyQRData.transactionId);
      message.success('Transaction ID 已複製到剪貼簿');
    }
  };

  const handleOpenDeepLink = () => {
    if (verifyQRData?.authUri) {
      window.open(verifyQRData.authUri, '_blank');
    }
  };

  return (
    <div className="verify-panel">
      <Card className="verify-card">
        <Title level={3}>
          <CheckCircleOutlined /> 驗證優惠券
        </Title>

        {/* 步驟指示器 */}
        <Steps current={currentStep} style={{ marginBottom: 24 }}>
          <Step title="產生 QR Code" icon={<QrcodeOutlined />} />
          <Step title="消費者掃描" icon={<ScanOutlined />} />
          <Step title="完成驗證" icon={<CheckCircleOutlined />} />
        </Steps>

        {/* 步驟 1: 產生驗證 QR Code */}
        {currentStep === 0 && (
          <div style={{ textAlign: 'center', padding: '40px 0' }}>
            <Paragraph type="secondary">
              請先產生驗證 QR Code,讓消費者掃描以提供優惠券資訊
            </Paragraph>
            <Button 
              type="primary" 
              size="large"
              icon={<QrcodeOutlined />}
              onClick={handleGenerateQR}
              loading={loading}
            >
              產生驗證 QR Code
            </Button>
          </div>
        )}

        {/* 步驟 2: 等待消費者掃描並計算 */}
        {currentStep === 1 && (
          <>
            <Alert
              message="等待消費者掃描 QR Code"
              description={
                <div>
                  <p>請消費者使用數位憑證皮夾 APP 掃描 QR Code</p>
                  <p>
                    <Text strong>Transaction ID: </Text>
                    <Text code copyable>{verifyQRData?.transactionId}</Text>
                  </p>
                </div>
              }
              type="info"
              showIcon
              style={{ marginBottom: 24 }}
              action={
                <Button size="small" onClick={() => setQrModalVisible(true)}>
                  顯示 QR Code
                </Button>
              }
            />

            <Form
              form={form}
              layout="vertical"
              onFinish={onFinish}
              initialValues={{
                original_amount: 2000,
              }}
            >
              <Form.Item
                label="消費金額 (新台幣)"
                name="original_amount"
                rules={[{ required: true, message: '請輸入消費金額' }]}
              >
                <InputNumber
                  style={{ width: '100%' }}
                  size="large"
                  min={0}
                  formatter={value => `$ ${value}`.replace(/\B(?=(\d{3})+(?!\d))/g, ',')}
                  parser={value => value.replace(/\$\s?|(,*)/g, '')}
                />
              </Form.Item>

              <Form.Item>
                <Space style={{ width: '100%' }}>
                  <Button 
                    type="primary" 
                    htmlType="submit" 
                    size="large"
                    loading={loading}
                    icon={<CheckCircleOutlined />}
                  >
                    驗證並計算折扣
                  </Button>
                  <Button 
                    size="large"
                    onClick={resetForm}
                  >
                    重新開始
                  </Button>
                </Space>
              </Form.Item>
            </Form>
          </>
        )}

        {/* 步驟 3: 顯示驗證結果 */}
        {currentStep === 2 && verifyResult && (
          <>
            <Alert
              message="驗證成功!"
              description={verifyResult.message}
              type="success"
              showIcon
              style={{ marginBottom: 24 }}
            />
            
            <Row gutter={16}>
              <Col span={8}>
                <Card>
                  <Statistic
                    title="消費者"
                    value={verifyResult.customer_name}
                    prefix={<TagOutlined />}
                  />
                </Card>
              </Col>
              <Col span={8}>
                <Card>
                  <Statistic
                    title="原價"
                    value={verifyResult.original_amount}
                    precision={0}
                    prefix="$"
                    suffix="元"
                  />
                </Card>
              </Col>
              <Col span={8}>
                <Card>
                  <Statistic
                    title="折扣金額"
                    value={verifyResult.discount_amount}
                    precision={0}
                    valueStyle={{ color: '#cf1322' }}
                    prefix="-$"
                    suffix="元"
                  />
                </Card>
              </Col>
            </Row>

            <Row gutter={16} style={{ marginTop: 16 }}>
              <Col span={12}>
                <Card>
                  <Statistic
                    title="到期日期"
                    value={verifyResult.expired_date}
                  />
                </Card>
              </Col>
              <Col span={12}>
                <Card style={{ background: '#f6ffed', borderColor: '#b7eb8f' }}>
                  <Statistic
                    title="實付金額"
                    value={verifyResult.final_amount}
                    precision={0}
                    valueStyle={{ color: '#3f8600', fontSize: '32px', fontWeight: 'bold' }}
                    prefix={<DollarOutlined />}
                    suffix="元"
                  />
                </Card>
              </Col>
            </Row>

            <Divider />

            <Space style={{ width: '100%', justifyContent: 'center' }}>
              <Button 
                type="primary"
                size="large"
                onClick={resetForm}
              >
                驗證下一位消費者
              </Button>
            </Space>
          </>
        )}
      </Card>

      {/* ⭐ QR Code 顯示 Modal - 修正版本 */}
      <Modal
        title={<><QrcodeOutlined /> 驗證 QR Code</>}
        open={qrModalVisible}
        onCancel={() => setQrModalVisible(false)}
        footer={[
          <Button key="copy" icon={<LinkOutlined />} onClick={handleCopyTransactionId}>
            複製 Transaction ID
          </Button>,
          <Button key="deeplink" onClick={handleOpenDeepLink}>
            開啟 APP
          </Button>,
          <Button key="close" type="primary" onClick={() => setQrModalVisible(false)}>
            關閉
          </Button>
        ]}
        width={500}
      >
        {verifyQRData && (
          <div style={{ textAlign: 'center' }}>
            {/* ⭐ 直接顯示圖片,不使用 QRCode 組件 */}
            <img 
              src={verifyQRData.qrcodeImage} 
              alt="Verify QR Code"
              style={{ 
                width: '300px', 
                height: '300px',
                margin: '20px auto',
                display: 'block',
                border: '1px solid #d9d9d9',
                borderRadius: '4px'
              }}
              onError={(e) => {
                console.error('QR Code image load error');
                e.target.style.display = 'none';
              }}
            />
            
            <Paragraph>
              <Text strong>Transaction ID:</Text><br />
              <Text copyable code>{verifyQRData.transactionId}</Text>
            </Paragraph>
            
            <Paragraph type="secondary">
              請消費者使用數位憑證皮夾 APP 掃描此 QR Code
            </Paragraph>
            
            {/* Deep Link 連結 */}
            {verifyQRData.authUri && (
              <Paragraph>
                <Text type="secondary" style={{ fontSize: '12px' }}>
                  或點擊下方按鈕直接開啟 APP
                </Text>
              </Paragraph>
            )}
          </div>
        )}
      </Modal>
    </div>
  );
}

export default VerifyPanel;