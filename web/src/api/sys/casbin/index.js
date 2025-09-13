import request from '@/utils/request';

// @Tags authority
// @Summary Update role API permissions
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateAuthorityPatams true "Update role API permissions"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Get Success"}"
// @Router /casbin/UpdateCasbin [post]
export async function UpdateCasbin(data) {
  const res = await request.post('/casbin/updateCasbin', data);
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

// @Tags casbin
// @Summary Get role API permission list
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateAuthorityPatams true "Get role API permission list"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Get Success"}"
// @Router /casbin/getPolicyPathByAuthorityId [post]
export async function getPolicyPathByAuthorityId(data) {
  const res = await request.post('/casbin/getPolicyPathByAuthorityId', data);
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}
