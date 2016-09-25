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

type JobsV3 struct {
	*RequestSchema
	/** Optional Job identifier */
	JobId *JobKeyV3 `json:"job_id"`

	/** jobs */
	Jobs []*JobV3 `json:"jobs"`

	/* INHERITED: Comma-separated list of JSON field paths to exclude from the result, used like: "/3/Frames?_exclude_fields=frames/frame_id/URL,__meta"
	ExcludeFields string: "" `json:"_exclude_fields"`
	*/
}

func NewJobsV3() *JobsV3 {
	return &JobsV3{
		JobId: nil,
		Jobs:  nil,
		RequestSchema: &RequestSchema{
			ExcludeFields: "",
		},
	}
}
