import axios from 'axios';
import qs from 'query-string';

export interface scopeRes {
  [key: string]: string[];
}
export function getRoles() {
  return axios.get('/admin/user/roles');
}
export function switchRole(param:any){
  return axios.post('/admin/user/switchRole',param);
}
export interface ApiRecord {
  title: string;
  method: string;
  path: string;
  scope: string;
  group: string;
  need_login: number;
  need_auth: number;
}
// 删除api
export function delApi(id:number){
  return axios.get('/admin/baseApi/delete', {
    params: { id },
    paramsSerializer: (obj) => {
      return qs.stringify(obj);
    },
  });
}
export function saveApi(param: ApiRecord) {
  return axios.post('/admin/baseApi/save', param);
}
export function delRole(id:number){
  return axios.get('/admin/role/delete', {
    params: { id },
    paramsSerializer: (obj) => {
      return qs.stringify(obj);
    },
  });
}
export interface ApiListRecord {
  [key: string]: {
    [key: string]: ApiRecord[];
  };
}
export interface ApiListRes {
  all_scope?: scopeRes;
  formatted_list?: ApiListRecord;
}
export function queryApiList() {
  return axios.get<ApiListRes>('/admin/baseApi/list');
}
// api列表相关

//  菜单列表相关
export interface MenuMeta {
  locale?: string;
  requiresAuth: boolean;
  icon?: string;
  hideInMenu?: boolean;
  hideChildrenInMenu?: boolean;
  order?: number;
}
export interface MenuListRecord {
  id?: number;
  name: string;
  title: string;
  path: string;
  type: number;
  parent_id?: number | null;
  meta: MenuMeta;
  children?: MenuListRecord[];
}
export interface MenuListRes {
  list: MenuListRecord[];
  total: number;
}
export function delMenu(id: number) {
  return axios.get<any>('/admin/menu/delete', {
    params: { id },
    paramsSerializer: (obj) => {
      return qs.stringify(obj);
    },
  });
}
export function queryMenuList() {
  return axios.get<MenuListRes>('/admin/menu/list');
}
export function saveMenu(params: MenuListRecord) {
  return axios.post('/admin/menu/save', params);
}
// 管理员列表相关
export interface AdminListRecord {
  id: number;
  accountId: string;
  name: string;
  avatar: string;
  email: string;
  account: string;
  organizationName: string;
  jobName: string;
  phone: string;
  status: number;
  registrationDate: string;
}

export interface AdminListRes {
  list: AdminListRecord[];
  total: number;
}
export interface AdminListParam {
  current: number;
  pageSize: number;
  name?: string;
  email?: string;
  nickname?: string;
  account?: string;
  mobile?: string;
}
export interface AdminLogParam {
  current: number;
  pageSize: number;
  nickname?:string;
  api_name?:string;
  path?:string;
}
export function delAdmin(id: number) {
  return axios.get<any>('/admin/user/delete', {
    params: { id },
    paramsSerializer: (obj) => {
      return qs.stringify(obj);
    },
  });
}
export function queryAdminList(params: AdminListParam) {
  return axios.get<AdminListRes>('/admin/user/list', {
    params,
    paramsSerializer: (obj) => {
      return qs.stringify(obj);
    },
  });
}
export function queryLogs(params: AdminListParam) {
  return axios.get<AdminLogParam>('/admin/adminLogs/list', {
    params,
    paramsSerializer: (obj) => {
      return qs.stringify(obj);
    },
  });
}
export function saveAdmin(params: any) {
  return axios.post('/admin/user/save', params);
}

// 角色列表相关
export interface RoleListRecord {
  id: number;
  name?: string;
  description?: string;
  created_time: string;
  updated_time: string;
}
export interface RoleListRes {
  list: RoleListRecord[];
  total: number;
}
export function saveRole(record:RoleListRecord){
  return axios.post('/admin/role/save', record)
}
export function setUserRole(params: any){
  return axios.post('/admin/role/setUserRole', params);
}
export function queryUserRole(params:any){
  return axios.get<RoleListRecord[]>('/admin/role/userRoles', {
    params: {
      uid:params.id
    },
    paramsSerializer: (obj) => {
      return qs.stringify(obj);
    },
  });
}
export function authRoleMenu(params: any) {
  return axios.post('/admin/role/accreditMenu', params);
}
export function authRoleApi(params: any) {
  return axios.post('/admin/role/accreditApi', params);
}
export function getRoleAccess(roleId: number) {
  return axios.get<any>('/admin/role/roleAccess', {
    params: { id: roleId },
    paramsSerializer: (obj) => {
      return qs.stringify(obj);
    },
  });
}
export interface RoleAccessRes {
  tree: {
    title: string;
    id?: number | string;
    children?: {
      title: string;
      id?: number | string;
      children?: {
        title: string;
        id?: number | string;
        path?: string;
      };
    }[];
  }[];
  menu_list: MenuListRecord[];
  api_list: {
    [key: string]: {
      [innerKey: string]: ApiRecord[];
    };
  };
}
export function getAuthNode() {
  return axios.get<RoleAccessRes>('/admin/role/authNode');
}
export interface RoleListParams {
  current: number;
  pageSize: number;
  name?: string;
}
export function queryRoleList(params: RoleListParams) {
  return axios.get<RoleListRes>('/admin/role/list', {
    params,
    paramsSerializer: (obj) => {
      return qs.stringify(obj);
    },
  });
}
