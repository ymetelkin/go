package appl

import (
	"strings"

	"github.com/ymetelkin/go/json"
)

type state struct {
	Code string
	Name string
}

func (st *state) ToJson() *json.Object {
	jo := json.Object{}
	jo.AddString("code", st.Code)
	jo.AddString("name", st.Name)
	jo.AddString("type", "AUDGEOGRAPHY")
	return &jo
}

func getState(s string) *state {
	if s == "" || len(s) < 2 {
		return nil
	}

	key := strings.ToUpper(s)
	if strings.Index(s, "NYC") == 0 {
		return &state{"b836d07082c610048807df092526b43e", "New York City"}
	} else {
		runes := []rune(key)
		key = string(runes[0:2])

		switch key {
		case "AL":
			return &state{"b8099e4881d610048a11df092526b43e", "Alabama"}
		case "AK":
			return &state{"cbb727a881d610048a29df092526b43e", "Alaska"}
		case "AZ":
			return &state{"e427079081d610048a4edf092526b43e", "Arizona"}
		case "AR":
			return &state{"687e74a082af1004823adf092526b43e", "Arkansas"}
		case "CA":
			return &state{"789fdd8882af10048263df092526b43e", "California"}
		case "CO":
			return &state{"902a5eb082af1004828adf092526b43e", "Colorado"}
		case "CT":
			return &state{"a42dc0a082af100482a7df092526b43e", "Connecticut"}
		case "DE":
			return &state{"bcadd4f882af100482c9df092526b43e", "Delaware"}
		case "FL":
			return &state{"cb06ab1082af100482f8df092526b43e", "Florida"}
		case "GA":
			return &state{"dec1cce882af10048320df092526b43e", "Georgia"}
		case "HI":
			return &state{"ee324cc082af10048342df092526b43e", "Hawaii"}
		case "ID":
			return &state{"1885b7f082b01004835edf092526b43e", "Idaho"}
		case "IL":
			return &state{"2c6a186082b010048379df092526b43e", "Illinois"}
		case "IN":
			return &state{"0760000082b2100483c7df092526b43e", "Indiana"}
		case "IA":
			return &state{"1608ba1082b310048433df092526b43e", "Iowa"}
		case "KS":
			return &state{"1e8c5a7082b310048450df092526b43e", "Kansas"}
		case "KY":
			return &state{"2f6e294082b310048474df092526b43e", "Kentucky"}
		case "LA":
			return &state{"43fb970882b310048496df092526b43e", "Louisiana"}
		case "ME":
			return &state{"8d2caa7082b3100484b8df092526b43e", "Maine"}
		case "MD":
			return &state{"b0fa317082b3100484dbdf092526b43e", "Maryland"}
		case "MA":
			return &state{"bed6942882b310048501df092526b43e", "Massachusetts"}
		case "MI":
			return &state{"6bf49b4082c410048696df092526b43e", "Michigan"}
		case "MN":
			return &state{"9f12355082c4100486addf092526b43e", "Minnesota"}
		case "MS":
			return &state{"b3dfffa882c4100486c3df092526b43e", "Mississippi"}
		case "MO":
			return &state{"bd8c35d082c4100486d5df092526b43e", "Missouri"}
		case "MT":
			return &state{"6429117882c610048770df092526b43e", "Montana"}
		case "NE":
			return &state{"808300b882c610048788df092526b43e", "Nebraska"}
		case "NV":
			return &state{"8bb89dd082c61004879fdf092526b43e", "Nevada"}
		case "NH":
			return &state{"9531546082c6100487b5df092526b43e", "New Hampshire"}
		case "NJ":
			return &state{"a0eed68882c6100487cddf092526b43e", "New Jersey"}
		case "NM":
			return &state{"aacce28082c6100487e4df092526b43e", "New Mexico"}
		case "NY":
			return &state{"b58f18a082c6100487fbdf092526b43e", "New York"}
		case "NC":
			return &state{"c01d179082c610048813df092526b43e", "North Carolina"}
		case "ND":
			return &state{"cbaeb75882c61004882adf092526b43e", "North Dakota"}
		case "OH":
			return &state{"dcb000c082c610048843df092526b43e", "Ohio"}
		case "OK":
			return &state{"f142e8e082c610048858df092526b43e", "Oklahoma"}
		case "OR":
			return &state{"fe016fe882c61004886adf092526b43e", "Oregon"}
		case "PA":
			return &state{"0b394d7082c71004887fdf092526b43e", "Pennsylvania"}
		case "RI":
			return &state{"1bf4bc0882c71004889cdf092526b43e", "Rhode Island"}
		case "SC":
			return &state{"29d11ec082c7100488aedf092526b43e", "South Carolina"}
		case "SD":
			return &state{"5578469882c7100488badf092526b43e", "South Dakota"}
		case "TN":
			return &state{"62532b5882c7100488cedf092526b43e", "Tennessee"}
		case "TX":
			return &state{"6e92d9b882c7100488e5df092526b43e", "Texas"}
		case "UT":
			return &state{"c1dff44882c710048903df092526b43e", "Utah"}
		case "VT":
			return &state{"d2f8d8a882c710048915df092526b43e", "Vermont"}
		case "VA":
			return &state{"eaed376082c71004892cdf092526b43e", "Virginia"}
		case "WA":
			return &state{"08a0a00882c810048942df092526b43e", "Washington"}
		case "WV":
			return &state{"130bcce882c81004895adf092526b43e", "West Virginia"}
		case "WI":
			return &state{"1bc1bc3082c81004896cdf092526b43e", "Wisconsin"}
		case "WY":
			return &state{"2fb83d4082c810048984df092526b43e", "Wyoming"}
		case "DC":
			return &state{"788b364882f110048df3df092526b43e", "District of Columbia"}
		}
	}
	return nil
}
