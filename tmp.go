// reader := csv.NewReader(file)
// var readerCount int
// key := make(map[string]int)
// for {
// 	record, err := reader.Read()
// 	if err == io.EOF {
// 		break
// 	}
// 	if err != nil {
// 		respond.With(w, r, http.StatusBadRequest, nil, err)
// 		return
// 	}
// 	if readerCount == 0 {
// 		for i, e := range record {
// 			key[e] = i
// 		}
// 	} else {
// 		id, err := uuid.FromString(record[key["id"]])
// 		if err != nil {
// 			respond.With(w, r, http.StatusBadRequest, nil, err)
// 		}
// 		c := &Contact{
// 			ID:        id,
// 			FirstName: record[key["first_name"]],
// 			LastName:  record[key["last_name"]],
// 			Email:     record[key["email"]],
// 			Phone:     record[key["phone"]],
// 		}
// 		if err := c.updateContact(db); err != nil {
// 			respond.With(w, r, http.StatusInternalServerError, nil, err.Error())
// 			return
// 		}
// 	}
// 	readerCount++
// 	// results = append(results, record)
// }
//
// respond.With(w, r, http.StatusOK, nil, nil)
