package accounts

import (
	"time"

	"golang.org/x/text/language"
    //"golang.org/x/text/language/display"
)
/*

Templates
Template Files
Template Tags
Signals
user_signed_up
user_sign_up_attempt
user_logged_in
user_login_attempt
signup_code_sent
signup_code_used
email_confirmed
email_confirmation_sent
password_changed
password_expired
Management Commands


Settings
ACCOUNT_OPEN_SIGNUP
Default: True

ACCOUNT_LOGIN_URL
Default: "account_login"

ACCOUNT_SIGNUP_REDIRECT_URL
Default: "/"

ACCOUNT_LOGIN_REDIRECT_URL
Default: "/"

ACCOUNT_LOGOUT_REDIRECT_URL
Default: "/"

ACCOUNT_PASSWORD_CHANGE_REDIRECT_URL
Default: "account_password"

ACCOUNT_PASSWORD_RESET_REDIRECT_URL
Default: "account_login"

ACCOUNT_PASSWORD_EXPIRY
Default: 0

ACCOUNT_PASSWORD_USE_HISTORY
Default: False

ACCOUNT_REMEMBER_ME_EXPIRY
Default: 60 * 60 * 24 * 365 * 10

ACCOUNT_USER_DISPLAY
Default: lambda user: user.username

ACCOUNT_CREATE_ON_SAVE
Default: True

ACCOUNT_EMAIL_UNIQUE
Default: True

ACCOUNT_EMAIL_CONFIRMATION_REQUIRED
Default: False

ACCOUNT_EMAIL_CONFIRMATION_EMAIL
Default: True

ACCOUNT_EMAIL_CONFIRMATION_EXPIRE_DAYS
Default: 3

ACCOUNT_EMAIL_CONFIRMATION_ANONYMOUS_REDIRECT_URL
Default: "account_login"

ACCOUNT_EMAIL_CONFIRMATION_AUTHENTICATED_REDIRECT_URL
Default: None

ACCOUNT_EMAIL_CONFIRMATION_URL
Default: "account_confirm_email"

ACCOUNT_SETTINGS_REDIRECT_URL
Default: "account_settings"

ACCOUNT_NOTIFY_ON_PASSWORD_CHANGE
Default: True

ACCOUNT_DELETION_MARK_CALLBACK
Default: "account.callbacks.account_delete_mark"

ACCOUNT_DELETION_EXPUNGE_CALLBACK
Default: "account.callbacks.account_delete_expunge"

ACCOUNT_DELETION_EXPUNGE_HOURS
Default: 48

ACCOUNT_HOOKSET
Default: "account.hooks.AccountDefaultHookSet"

This setting allows you define your own hooks for specific functionality that django-user-accounts exposes. Point this to a class using a string and you can override the following methods:

send_invitation_email(to, ctx)
send_confirmation_email(to, ctx)
send_password_change_email(to, ctx)
send_password_reset_email(to, ctx)
ACCOUNT_TIMEZONES
Default: list(zip(pytz.all_timezones, pytz.all_timezones))

ACCOUNT_LANGUAGES

 */

type HookSet interface {

}

type AccountHooks struct {

}

var (
	ACCOUNT_OPEN_SIGNUP bool
	ACCOUNT_LOGIN_URL,
	ACCOUNT_SIGNUP_REDIRECT_URL,
	ACCOUNT_LOGIN_REDIRECT_URL,
	ACCOUNT_LOGOUT_REDIRECT_URL,
	ACCOUNT_PASSWORD_CHANGE_REDIRECT_URL,
	ACCOUNT_PASSWORD_RESET_REDIRECT_URL string

	ACCOUNT_PASSWORD_EXPIRY int
	ACCOUNT_PASSWORD_USE_HISTORY bool
	ACCOUNT_REMEMBER_ME_EXPIRY time.Duration
	ACCOUNT_USER_DISPLAY string
	ACCOUNT_CREATE_ON_SAVE bool
	ACCOUNT_EMAIL_UNIQUE bool
	ACCOUNT_EMAIL_CONFIRMATION_REQUIRED bool
	ACCOUNT_EMAIL_CONFIRMATION_EMAIL bool
	ACCOUNT_EMAIL_CONFIRMATION_EXPIRE_DAYS int
	ACCOUNT_EMAIL_CONFIRMATION_ANONYMOUS_REDIRECT_URL,
	ACCOUNT_EMAIL_CONFIRMATION_AUTHENTICATED_REDIRECT_URL,
	ACCOUNT_EMAIL_CONFIRMATION_URL,
	ACCOUNT_SETTINGS_REDIRECT_URL string
	ACCOUNT_NOTIFY_ON_PASSWORD_CHANGE bool
	ACCOUNT_DELETION_MARK_CALLBACK bool
	ACCOUNT_DELETION_EXPUNGE_CALLBACK string
	ACCOUNT_DELETION_EXPUNGE_HOURS int
	ACCOUNT_HOOKSET  HookSet
	ACCOUNT_TIMEZONES string
	ACCOUNT_LANGUAGES []language.Tag


)
// Settings



func init(){
	ACCOUNT_OPEN_SIGNUP = true

	ACCOUNT_LOGIN_URL = "/login"


	ACCOUNT_SIGNUP_REDIRECT_URL = "/"


	ACCOUNT_LOGIN_REDIRECT_URL = "/"


	ACCOUNT_LOGOUT_REDIRECT_URL = "/"


	ACCOUNT_PASSWORD_CHANGE_REDIRECT_URL = "/"


	ACCOUNT_PASSWORD_RESET_REDIRECT_URL = "/login"


	ACCOUNT_PASSWORD_EXPIRY = 0


	ACCOUNT_PASSWORD_USE_HISTORY = false


	ACCOUNT_REMEMBER_ME_EXPIRY = time.Duration(60 * 60 * 24 * 365 * 10)*time.Second



	ACCOUNT_USER_DISPLAY =""


	ACCOUNT_CREATE_ON_SAVE = true


	ACCOUNT_EMAIL_UNIQUE = true


	ACCOUNT_EMAIL_CONFIRMATION_REQUIRED = false


	ACCOUNT_EMAIL_CONFIRMATION_EMAIL = true


	ACCOUNT_EMAIL_CONFIRMATION_EXPIRE_DAYS = 3


	ACCOUNT_EMAIL_CONFIRMATION_ANONYMOUS_REDIRECT_URL = "/login"


	ACCOUNT_EMAIL_CONFIRMATION_AUTHENTICATED_REDIRECT_URL = "/"


	ACCOUNT_EMAIL_CONFIRMATION_URL = "account_confirm_email"


	ACCOUNT_SETTINGS_REDIRECT_URL = "account_settings"


	ACCOUNT_NOTIFY_ON_PASSWORD_CHANGE = true


	ACCOUNT_DELETION_MARK_CALLBACK = "delete"

	ACCOUNT_DELETION_EXPUNGE_CALLBACK = "expunge"

	ACCOUNT_DELETION_EXPUNGE_HOURS = 48

	ACCOUNT_HOOKSET = AccountHooks{}



	ACCOUNT_TIMEZONES = ""

	ACCOUNT_LANGUAGES = []language.Tag{
		language.AmericanEnglish, // en-US fallback
		language.German,          // de
	}


}



