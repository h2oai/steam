/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package bindings

type TwoDimTableV3 struct {
	*TwoDimTableBase
	/* INHERITED: Table Name
	Name string: "" `json:"name"`
	*/

	/* INHERITED: Table Description
	Description string: "" `json:"description"`
	*/

	/* INHERITED: Column Specification
	Columns []*ColumnSpecsBase: nil `json:"columns"`
	*/

	/* INHERITED: Number of Rows
	Rowcount int32: 0 `json:"rowcount"`
	*/

	/* INHERITED: Table Data (col-major)
	Data [][]Polymorphic: nil `json:"data"`
	*/
}

func NewTwoDimTableV3() *TwoDimTableV3 {
	return &TwoDimTableV3{
		TwoDimTableBase: &TwoDimTableBase{
			Name:        "",
			Description: "",
			Columns:     nil,
			Rowcount:    0,
			Data:        nil,
		},
	}
}
