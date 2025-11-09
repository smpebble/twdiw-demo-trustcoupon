import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080/api';

const api = {
  // 發行端 API
  issueCoupon: (data) => {
    return axios.post(`${API_BASE_URL}/issue`, data);
  },

  getTransaction: (transactionId) => {
    return axios.get(`${API_BASE_URL}/transaction/${transactionId}`);
  },

  // ⭐ 驗證端 API - 修正
  generateVerifyQR: () => {
    // ⭐ 不需要參數,直接呼叫
    return axios.post(`${API_BASE_URL}/verify/qrcode`);
  },

  getVerifyResult: (transactionId) => {
    return axios.post(`${API_BASE_URL}/verify/result`, { 
      transaction_id: transactionId 
    });
  },

  calculateDiscount: (data) => {
    return axios.post(`${API_BASE_URL}/verify/calculate`, data);
  },
};

export default api;