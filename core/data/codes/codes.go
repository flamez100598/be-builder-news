package codes

type Code int

const (
	Ok Code = 1000
	NotOk Code = 1001
	Yes Code  = 1002
	No Code = 1003
	Success Code = 1004
	Fail Code = 1005
	Right Code = 1006
	Wrong Code = 1007
	NotFound Code = 1008
	InternalError Code = 1009
	Existed Code = 1010
	NoPermission Code = 1011
	FileTooLarge Code = 1012
	Deleted Code = 1013
	UsernamePasswordIncorrect Code = 1100
)

func (this Code) Value() int {
	return int(this)
}

func (this Code) ToInt32() int32 {
	return int32(this)
}

func (this Code) Msg() string {
	switch this {
	case Ok:
		return "Ok"
	case NotOk:
		return "Not Ok"
	case Yes:
		return "Yes"
	case No:
		return "No"
	case Success:
		return "Success"
	case Fail:
		return "Fail"
	case Right:
		return "Right"
	case Wrong:
		return "Wrong"
	case NotFound:
		return "Not Found"
	case InternalError:
		return "Internal Error"
	case Existed:
		return "Existed"
	case NoPermission:
		return "No Permission"
	case FileTooLarge:
		return "File Too Large"
	case UsernamePasswordIncorrect:
		return "Username or Password Incorrect"
	case Deleted:
		return "Deleted"
	default:
		return "Unknown"
	}
}