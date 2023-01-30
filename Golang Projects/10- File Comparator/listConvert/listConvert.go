package listConvert

func ListConvert(complexList [][]string, basicList *[]string) {
	for _, i := range complexList {
		for _, l := range i {
			*basicList = append(*basicList, l)
		}
	}
}
