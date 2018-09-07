/*
*Desc:
*CreateBy:Cooyw
*Time:2018/9/6
*/
package mydatafile

import (
	"os"
	"sync"
	"github.com/pkg/errors"
)

//用于表示数据文件的接口类型
type DataFile interface {
	//读取一个数据块
	Read() (rsn int64, d Data, err error)
	//写一个数据块
	Write(d Data) (wsn int64, err error)
	//获取最后读取的数据块序列号
	RSN() int64
	//获取最后写入的数据块序列号
	WSN() int64
	//获取数据块的长度
	DataLen() uint32
	//关闭数据文件
	Close() error
}

//用于表示数据的类型
type Data []byte

//用于标书数据文件的实现类型
type myDataFile struct {
	f       *os.File     //	文件
	fmutex  sync.RWMutex //	用于文件的读写锁
	woffset int64        //	写操作需要用到的偏移量
	roffset int64        //	读操作需要用到的偏移量
	wmutex  sync.Mutex   //	写操作需要用到的互斥锁
	rmutex  sync.Mutex   //	读操作需要用到的互斥锁
	dataLen uint32       //	数据块长度
}

func NewDataFile(path string, dataLen uint32) (DataFile, error) {

	f, err := os.Create(path)
	if err != nil {
		//exist := os.IsExist(err)
		return nil, err
	}
	if dataLen == 0 {
		return nil, errors.New("Invalid data length!")
	}
	df := &myDataFile{f: f, dataLen: dataLen}

	return df, nil

}

func (df *myDataFile) Read() (rsn int64, d Data, err error) {
	//读取并更新读偏移量
	var offset int64
	df.rmutex.Lock()
	offset = df.roffset
	df.roffset += int64(df.dataLen)
	df.rmutex.Unlock()

	//读取一个数据块
	rsn = offset / int64(df.dataLen)
	df.fmutex.RLock()
	defer df.fmutex.RUnlock()

	bytes := make([]byte, df.dataLen)
	_, err = df.f.ReadAt(bytes,offset)
	if err != nil {
		return
	}
	d = bytes
	return
}

//func (df *myDataFile) Read() (rsn int64, d Data, err error) {
//	//读取并更新读偏移量
//	var offset int64
//	df.rmutex.Lock()
//	offset = df.roffset
//	df.roffset += int64(df.dataLen)
//	df.rmutex.Unlock()
//
//	//读取一个数据块
//	rsn = offset / int64(df.dataLen)
//	bytes := make([]byte, df.dataLen)
//	for {
//		df.fmutex.RLock()
//		_, err = df.f.ReadAt(bytes, offset)
//		if err != nil {
//			if err == io.EOF {
//				df.fmutex.RUnlock()
//				continue
//			}
//			df.fmutex.RUnlock()
//			return
//		}
//		d = bytes
//		df.fmutex.RUnlock()
//		return
//	}
//}

func (df *myDataFile) Write(d Data) (wsn int64, err error) {
	//读取并更新偏移量
	var offset int64
	df.wmutex.Lock()
	offset = df.woffset
	df.woffset += int64(df.dataLen)
	df.wmutex.Unlock()

	//写入一个数据块
	wsn = offset / int64(df.dataLen)
	var bytes Data
	if len(d) > int(df.dataLen) {
		bytes = d[0:df.dataLen]
	} else {
		bytes = d
	}
	df.fmutex.Lock()
	defer df.fmutex.Unlock()
	_, err = df.f.Write(bytes)
	return
}

func (df *myDataFile) RSN() int64 {
	df.rmutex.Lock()
	defer df.rmutex.Unlock()
	return df.roffset / int64(df.dataLen)
}

func (df *myDataFile) WSN() int64 {
	df.wmutex.Lock()
	defer df.wmutex.Unlock()
	return df.woffset / int64(df.dataLen)
}

func (df *myDataFile) DataLen() uint32 {
	return df.dataLen
}

func (df *myDataFile) Close() error {
	return df.f.Close()
}
