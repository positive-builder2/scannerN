package scanner

func GuessService(port int) string {
	switch port {
	case 22:
		return "ssh"
	case 80:
		return "http"
	case 443:
		return "https"
	case 21:
		return "ftp"
	case 25:
		return "smtp"
	case 3306:
		return "mysql"
	case 5432:
		return "postgres"
	default:
		return "unknown"
	}
}
