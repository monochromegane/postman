package postman

type Option struct {
	Dir string `short:"d" description:"starts watching the named directory." default:"."`
}
