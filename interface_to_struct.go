type RepoStruct struct {
	sdk.GetRepoOutput
}

func (s *RepoStruct) FillStruct(i interface{}) error {
	m, ok := reflect.ValueOf(i).Interface().(map[string]interface{})
	if !ok {
		return errors.New("FillStruct conver interface to map[string]interface{} error")
	}
	for k, v := range m {
		err := SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}


rootNodeData := &RepoStruct{}
err = rootNodeData.FillStruct(rootNode.Data)
  
  
