package generic

type DataInterface interface {
	GetData() int64
}

type DataType1 struct {
	data int64
}

func (cls *DataType1) GetData() int64 {
	return cls.data
}

type DataType2 struct {
	data int64
}

func (cls *DataType2) GetData() int64 {

	return cls.data
}

func GetMyData(dataType int64) DataInterface {
	if dataType == 1 {
		return new(DataType1)
	} else {
		return new(DataType2)
	}
}
