import React, { useState } from 'react';
import { Layout, Tabs, message } from 'antd';
import IssuePanel from './components/IssuePanel';
import VerifyPanel from './components/VerifyPanel';
import 'antd/dist/reset.css';
import './App.css';

const { Header, Content } = Layout;
const { TabPane } = Tabs;

function App() {
  const [activeTab, setActiveTab] = useState('issue');

  return (
    <Layout className="app-layout">
      <Header className="app-header">
        <div className="header-content">
          <h1>🎫 TrustCoupon 信任券鏈系統</h1>
          <p className="merchant-name">商家:一路發發</p>
        </div>
      </Header>
      
      <Content className="app-content">
        <div className="content-wrapper">
          <Tabs 
            activeKey={activeTab} 
            onChange={setActiveTab}
            size="large"
            centered
          >
            <TabPane tab="📤 發行優惠券" key="issue">
              <IssuePanel />
            </TabPane>
            <TabPane tab="✅ 驗證優惠券" key="verify">
              <VerifyPanel />
            </TabPane>
          </Tabs>
        </div>
      </Content>
    </Layout>
  );
}

export default App;