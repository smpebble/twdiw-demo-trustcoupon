import React, { useState } from 'react';
import { Button, Modal } from 'antd';
import { ThunderboltOutlined } from '@ant-design/icons';
import './SpinWheel.css';

const SpinWheel = ({ onResult }) => {
  const [isSpinning, setIsSpinning] = useState(false);
  const [rotation, setRotation] = useState(0);
  const [modalVisible, setModalVisible] = useState(false);
  const [result, setResult] = useState(null);

  // 折扣金額選項 (100-900)
  const amounts = [100, 200, 300, 400, 500, 600, 700, 800, 900];
  const colors = [
    '#FF6B6B', '#4ECDC4', '#45B7D1', '#FFA07A',
    '#98D8C8', '#F7DC6F', '#BB8FCE', '#85C1E2', '#F8B739'
  ];

  const handleSpin = () => {
    if (isSpinning) return;

    setIsSpinning(true);
    setModalVisible(true);

    // 隨機選擇一個金額
    const randomIndex = Math.floor(Math.random() * amounts.length);
    const selectedAmount = amounts[randomIndex];

    // 計算旋轉角度 (每個扇形 40 度)
    const degreesPerSection = 360 / amounts.length;
    const targetDegree = randomIndex * degreesPerSection;
    
    // 多轉幾圈 + 目標角度
    const extraSpins = 5; // 額外轉 5 圈
    const totalRotation = rotation + (360 * extraSpins) + (360 - targetDegree);

    setRotation(totalRotation);
    setResult(selectedAmount);

    // 3秒後停止
    setTimeout(() => {
      setIsSpinning(false);
      onResult(selectedAmount);
    }, 3000);
  };

  const closeModal = () => {
    setModalVisible(false);
  };

  return (
    <>
      <Button
        type="dashed"
        icon={<ThunderboltOutlined />}
        onClick={handleSpin}
        disabled={isSpinning}
        size="large"
      >
        幸運輪盤
      </Button>

      <Modal
        title="🎰 幸運折扣輪盤"
        open={modalVisible}
        onCancel={closeModal}
        footer={[
          <Button key="close" type="primary" onClick={closeModal} disabled={isSpinning}>
            {isSpinning ? '旋轉中...' : '確定'}
          </Button>
        ]}
        width={600}
        centered
      >
        <div className="spin-wheel-container">
          {/* 輪盤主體 */}
          <div className="wheel-wrapper">
            <div
              className="wheel"
              style={{
                transform: `rotate(${rotation}deg)`,
                transition: isSpinning ? 'transform 3s cubic-bezier(0.17, 0.67, 0.12, 0.99)' : 'none'
              }}
            >
              {amounts.map((amount, index) => {
                const rotation = (360 / amounts.length) * index;
                return (
                  <div
                    key={amount}
                    className="wheel-section"
                    style={{
                      transform: `rotate(${rotation}deg)`,
                      backgroundColor: colors[index]
                    }}
                  >
                    <div className="wheel-text">
                      ${amount}
                    </div>
                  </div>
                );
              })}
            </div>

            {/* 中心圓點 */}
            <div className="wheel-center">
              <ThunderboltOutlined style={{ fontSize: '32px', color: '#fff' }} />
            </div>

            {/* 指針 */}
            <div className="wheel-pointer"></div>
          </div>

          {/* 結果顯示 */}
          {!isSpinning && result && (
            <div className="result-display">
              <h2>🎉 恭喜!</h2>
              <p>折扣金額: <span className="result-amount">${result}</span></p>
              <p className="result-hint">金額已自動填入表單</p>
            </div>
          )}

          {isSpinning && (
            <div className="spinning-text">
              <p>🎰 旋轉中...</p>
            </div>
          )}
        </div>
      </Modal>
    </>
  );
};

export default SpinWheel;