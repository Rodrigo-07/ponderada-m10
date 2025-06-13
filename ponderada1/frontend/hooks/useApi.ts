import axios from 'axios';

export const api = axios.create({
  baseURL: 'http://10.128.0.245:8080/api/v1',
});
