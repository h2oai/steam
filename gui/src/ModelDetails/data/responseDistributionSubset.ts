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

export const responseDistributionSubset = [
    {
        "column": "tenure",
        "type": "numeric",
        "responseCounts": [
            {
                "churnedYesCount": 493,
                "churnedNoCount": 443,
                "value": "0 to 5"
            },
            {
                "churnedYesCount": 172,
                "churnedNoCount": 277,
                "value": "5 to 10"
            },
            {
                "churnedYesCount": 145,
                "churnedNoCount": 259,
                "value": "10 to 15"
            },
            {
                "churnedYesCount": 96,
                "churnedNoCount": 227,
                "value": "15 to 20"
            },
            {
                "churnedYesCount": 83,
                "churnedNoCount": 227,
                "value": "20 to 25"
            },
            {
                "churnedYesCount": 59,
                "churnedNoCount": 213,
                "value": "25 to 30"
            },
            {
                "churnedYesCount": 55,
                "churnedNoCount": 200,
                "value": "30 to 35"
            },
            {
                "churnedYesCount": 43,
                "churnedNoCount": 190,
                "value": "35 to 40"
            },
            {
                "churnedYesCount": 44,
                "churnedNoCount": 188,
                "value": "40 to 45"
            },
            {
                "churnedYesCount": 37,
                "churnedNoCount": 203,
                "value": "45 to 50"
            },
            {
                "churnedYesCount": 33,
                "churnedNoCount": 222,
                "value": "50 to 55"
            },
            {
                "churnedYesCount": 32,
                "churnedNoCount": 233,
                "value": "55 to 60"
            },
            {
                "churnedYesCount": 21,
                "churnedNoCount": 253,
                "value": "60 to 65"
            },
            {
                "churnedYesCount": 37,
                "churnedNoCount": 302,
                "value": "65 to 70"
            },
            {
                "churnedYesCount": 19,
                "churnedNoCount": 484,
                "value": "70 to 72"
            }
        ]
    },
    {
        "column": "gender",
        "type": "categorical",
        "responseCounts": [
            {
                "value": "Female",
                "churnedYesCount": 684,
                "churnedNoCount": 1934
            },
            {
                "value": "Male",
                "churnedYesCount": 685,
                "churnedNoCount": 1987
            }
        ]
    },
    {
        "column": "PhoneService",
        "type": "categorical",
        "responseCounts": [
            {
                "value": "No",
                "churnedYesCount": 117,
                "churnedNoCount": 389
            },
            {
                "value": "Yes",
                "churnedYesCount": 1252,
                "churnedNoCount": 3532
            }
        ]
    },
    {
        "column": "OnlineSecurity",
        "type": "categorical",
        "responseCounts": [
            {
                "value": "No",
                "churnedYesCount": 1068,
                "churnedNoCount": 1541
            },
            {
                "value": "Yes",
                "churnedYesCount": 222,
                "churnedNoCount": 1302
            },
            {
                "value": "No internet service",
                "churnedYesCount": 79,
                "churnedNoCount": 1078
            }
        ]
    },
    {
        "column": "Dependents",
        "type": "categorical",
        "responseCounts": [
            {
                "value": "No",
                "churnedYesCount": 1134,
                "churnedNoCount": 2567
            },
            {
                "value": "Yes",
                "churnedYesCount": 235,
                "churnedNoCount": 1354
            }
        ]
    }
];
