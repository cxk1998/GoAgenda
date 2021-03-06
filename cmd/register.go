package cmd

import (
	"github.com/chenf99/GoAgenda/entity"
	"github.com/chenf99/GoAgenda/service"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "用户注册",
	Long: `
GoAgenda register -u username -p password -e email -t telephone 
各个参数分别对应用户名、密码、邮箱地址、电话号码`,
	Run: func(cmd *cobra.Command, args []string) {
		// 获取参数值
		username, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("pass")
		email, _ := cmd.Flags().GetString("email")
		telephone, _ := cmd.Flags().GetString("tel")
		/*
		 * 参数格式、逻辑处理
		 * 1. 登陆与否判断
		 * 2. 参数格式合法性判断
		 * 3. 参数逻辑合法性判断
		 */
		// 1.登陆与否判断
		has_login := entity.CurStatus.GetStatus().Islogin
		// 已经登陆无法进行注册命令
		if has_login {
			service.Error.Println("GoAgenda " + username + "  register failed: You had already logined!")
			return
		}
		// 2. 参数格式合法性判断
		if username == "" {
			service.Error.Println("GoAgenda " + username + "  register failed: username cannot be null")
			return
		}
		if password == "" {
			service.Error.Println("GoAgenda " + username + "  register failed: password cannot be null")
			return
		}
		if email == "" {
			service.Error.Println("GoAgenda " + username + "  register failed: email cannot be null")
			return
		}
		if telephone == "" {
			service.Error.Println("GoAgenda " + username + "  register failed: telephone cannot be null")
			return
		}
		// 3. 参数逻辑合法性判断
		// 注册用户名不允许重复
		if service.UserModel.IsExist(username) {
			service.Error.Println("GoAgenda " + username + "  register failed: username had been existed!")
			return
		}

		/*
		 * 参数格式、逻辑合法后的响应处理
		 * 1. users.json添加一个用户
		 * 2. 写入日志并UI提示
		 */		
		service.Info.Println("GoAgenda " + username + "  register succeed!")

		userinfo := entity.UserData{
			Name : username,
			Password : password,
			Email : email,
			Tel : telephone,
		}
		service.UserModel.AddUser(userinfo)
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringP("user", "u", "", "your username")
	registerCmd.Flags().StringP("pass", "p", "", "your password")
	registerCmd.Flags().StringP("email", "e", "", "your email URL")
	registerCmd.Flags().StringP("tel", "t", "", "your telephone number")
}
