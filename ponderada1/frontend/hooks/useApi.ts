import axios from 'axios';

export const api = axios.create({
  baseURL: 'http://10.254.17.184:8080/api/v1',
});
