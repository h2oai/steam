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

type FrameV3 struct {
	*FrameBase
	/** Row offset to display */
	RowOffset int64 `json:"row_offset"`

	/** Number of rows to display */
	RowCount int32 `json:"row_count"`

	/** Column offset to return */
	ColumnOffset int32 `json:"column_offset"`

	/** Number of columns to return */
	ColumnCount int32 `json:"column_count"`

	/** Total number of columns in the Frame */
	TotalColumnCount int32 `json:"total_column_count"`

	/** checksum */
	Checksum int64 `json:"checksum"`

	/** Number of rows in the Frame */
	Rows int64 `json:"rows"`

	/** Number of columns in the Frame */
	NumColumns int64 `json:"num_columns"`

	/** Default percentiles, from 0 to 1 */
	DefaultPercentiles []float64 `json:"default_percentiles"`

	/** Columns in the Frame */
	Columns []*ColV3 `json:"columns"`

	/** Compatible models, if requested */
	CompatibleModels []string `json:"compatible_models"`

	/** Chunk summary */
	ChunkSummary *TwoDimTableV3 `json:"chunk_summary"`

	/** Distribution summary */
	DistributionSummary *TwoDimTableV3 `json:"distribution_summary"`

	/* INHERITED: Frame ID
	FrameId *FrameKeyV3: nil `json:"frame_id"`
	*/

	/* INHERITED: Total data size in bytes
	ByteSize int64: 0 `json:"byte_size"`
	*/

	/* INHERITED: Is this Frame raw unparsed data?
	IsText bool: false `json:"is_text"`
	*/
}

func NewFrameV3() *FrameV3 {
	return &FrameV3{
		RowOffset:           0,
		RowCount:            0,
		ColumnOffset:        0,
		ColumnCount:         0,
		TotalColumnCount:    0,
		Checksum:            0,
		Rows:                0,
		NumColumns:          0,
		DefaultPercentiles:  nil,
		Columns:             nil,
		CompatibleModels:    nil,
		ChunkSummary:        nil,
		DistributionSummary: nil,
		FrameBase: &FrameBase{
			FrameId:  nil,
			ByteSize: 0,
			IsText:   false,
		},
	}
}
