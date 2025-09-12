import request from '@/utils/request';

// @Router /authority/getAuthorityList [post]
export async function getAuthorityList(data) {
  const res = await request.post('/authority/getAuthorityList', data);
  if (res.data.code === 0 && res.data.data) {
    return res.data.data.list;
  }
  return Promise.reject(new Error(res.data.msg));
}

// @Summary Deleting a Role
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body {authorityId uint} true "Deleting a Role"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Get Success"}"
// @Router /authority/deleteAuthority [post]
export async function deleteAuthority(data) {
  const res = await request.post('/authority/deleteAuthority', data);
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// @Summary Create a character
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateAuthorityPatams true "Create a character"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Get Success"}"
// @Router /authority/createAuthority [post]
export async function createAuthority(data) {
  const res = await request.post('/authority/createAuthority', data);
  if (res.data.code === 0 && res.data.data) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// @Tags authority
// @Summary Copy a Role
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateAuthorityPatams true "Copy a Role"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Get Success"}"
// @Router /authority/copyAuthority [post]
export async function copyAuthority(data) {
  const res = await request.post('/authority/copyAuthority', data);
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// @Summary Setting Role Resource Permissions
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sysModel.SysAuthority true "Setting Role Resource Permissions"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Get Success"}"
// @Router /authority/setDataAuthority [post]
export async function setDataAuthority(data) {
  const res = await request.post('/authority/setDataAuthority', data);
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// @Summary Modify Role
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "Modify Role"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Get Success"}"
// @Router /authority/setDataAuthority [post]
export async function updateAuthority(data) {
  const res = await request.put('/authority/updateAuthority', data);
  if (res.data.code === 0 && res.data.data) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}
