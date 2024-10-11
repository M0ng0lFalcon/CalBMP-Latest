package FileUtil

import (
	"io"
	"log"
)

func MergeInpFile(inp1Path string, inp2Path string, targetPath string) {
	// Open inp file 1
	inp1File := OpenFileAsRead(inp1Path)

	// Open inp file 2
	inp2File := OpenFileAsRead(inp2Path)

	// Open output file
	out := OpenFileAsWrite(targetPath)

	_, err := io.Copy(out, inp1File)
	if err != nil {
		log.Fatalln("failed to append zip file to output:", err)
	}
	//log.Printf("wrote %d bytes of %s to output\n", n, inp1Path)

	_, err = io.Copy(out, inp2File)
	if err != nil {
		log.Fatalln("failed to append signed file to output:", err)
	}
	//log.Printf("wrote %d bytes of %s to output\n", n, inp2Path)

	// close file pointer
	_ = inp1File.Close()
	_ = inp2File.Close()
	_ = out.Close()

	// delete inp1 file and inp2 file
	//DeleteInp1Err := os.Remove(inp1Path)
	//if DeleteInp1Err != nil {
	//	log.Println(DeleteInp1Err)
	//}
	//DeleteInp2Err := os.Remove(inp2Path)
	//if DeleteInp2Err != nil {
	//	log.Println(DeleteInp2Err)
	//}
}
