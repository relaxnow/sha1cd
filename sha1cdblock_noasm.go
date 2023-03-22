package sha1cd

func block(dig *digest, p []byte) {
	blockGeneric(dig, p)
}
