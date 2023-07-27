package cft

import (
	"cl/internal/structs"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"io"
)

func (p *provider) GetMerchantList() (merchantList []structs.Merchant, err error) {
	var (
		pClientData      driver.Rows
		errorCode        int
		errorDescription string
	)

	// Execute CFT procedure
	_, err = p.cft.Exec(`BEGIN
		Z$MC_CLIENTS_HM_LIB_STAT_DATA.GETMERCHANT(
			DATA 				=> :DATA,
			ERRORCODE 			=> :ERRORCODE,
			ERRORDESCRIPTION 	=> :ERRORDESCRIPTION);
		END;`,

		// Extract data from response
		sql.Out{Dest: &pClientData},
		sql.Out{Dest: &errorCode},
		sql.Out{Dest: &errorDescription},
	)
	// Does not allow to cause panic
	defer pClientData.Close()

	if err != nil {
		p.logger.Error(err)
		return
	}

	if errorCode != 0 {
		err = errors.New(errorDescription)
		p.logger.Error(errorDescription)
		return
	}

	// log.Println(pClientData.Columns())

	//[MERCH_ID MERCHANT CATEGORY CITY]
	rows := make([]driver.Value, len(pClientData.Columns()))
	oneMerchant := structs.Merchant{}
	for {
		if err = pClientData.Next(rows); err != nil {
			if err == io.EOF {
				break
			}
			return
		}
		// log.Println(rows)
		if rows[1].(string) == oneMerchant.Name {
			oneMerchant.City = append(oneMerchant.City, rows[3].(string))
			continue
		}
		oneMerchant.MerchantID = fmt.Sprint(rows[0])
		oneMerchant.Name = rows[1].(string)
		oneMerchant.Category = rows[2].(string)
		oneMerchant.City = append(oneMerchant.City, rows[3].(string))
		// log.Println(oneMerchant)
		merchantList = append(merchantList, oneMerchant)
		oneMerchant = structs.Merchant{}
	}
	// log.Println(merchantList)

	// Debug mode
	// fmt.Println("==>", pClient)
	// fmt.Println("==>", pClientData.Columns())
	err = nil

	return
}
