const ApiBaseUrl = "/api/v1/";

export default {
  getLoginUrl: `${ApiBaseUrl}user/login`,
  getUserInfo: `${ApiBaseUrl}user`,
  getUserList: `${ApiBaseUrl}user/list`,
  checkUsername: `${ApiBaseUrl}user/check_username`,
  addUser: `${ApiBaseUrl}user/add_user`,
  deleteUser: `${ApiBaseUrl}user/delete_user`,
  deleteUserList: `${ApiBaseUrl}user/delete_user_list`,
  launchUser: `${ApiBaseUrl}user/recover_user`,
  infoBox: `${ApiBaseUrl}dashboard`
};
