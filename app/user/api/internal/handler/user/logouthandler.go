// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.1

package user

import (
	"net/http"

	"lucid/app/user/api/internal/logic/user"
	"lucid/app/user/api/internal/svc"

	"lucid/common/utils/response"
)

// 用户登出
func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewLogoutLogic(r.Context(), svcCtx)
		err := l.Logout()
		if err != nil {
			response.LogicError(r.Context(), w, err)
		} else {
			response.Ok(r.Context(), w, nil)
		}
	}
}
