import axios from 'axios';
import type { RouteRecordNormalized } from 'vue-router';
import { UserState } from '@/store/modules/user/types';

export interface LoginData {
  account: string;
  password: string;
}

export interface LoginRes {
  token: string;
}
export function login(data: LoginData) {
  return axios.post<LoginRes>('/admin/user/login', data);
}

export function logout() {
  return axios.get<LoginRes>('/admin/user/logout');
}

export function getUserInfo() {
  return axios.get<UserState>('/admin/user/info');
}

export function getMenuList() {
  return axios.get<RouteRecordNormalized[]>('/admin/user/authMenu');
}
