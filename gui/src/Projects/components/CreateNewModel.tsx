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

/**
 * Created by Jeff Fohl <jfohl@h2o.ai> on 7/11/16.
 */
 import * as React from 'react';
 import * as classNames from 'classnames';
 import PageHeader from './PageHeader';
 import Table from './Table';
 import Row from './Row';
 import Cell from './Cell';
 import ModelParameters from './ModelParameters';
 import { Link } from 'react-router';
 import '../styles/createnew.scss';

 interface Props {
   items: any[]
 }

 interface DispatchProps {
 }

 export default class CreateNewModel extends React.Component<Props & DispatchProps, any> {
   constructor() {
     super();
   }

   handleSubmit(e) {
     e.preventDefault();
   }

   render(): React.ReactElement<HTMLDivElement> {
     return (
       <div className="create-new-model">
         <PageHeader>
           Create New Model
           <div className="subtitle">Forked from <Link to={"/projects/models/DRF-1069085"}>DRF-1069085</Link></div>
         </PageHeader>
         <form name="new-model" onSubmit={this.handleSubmit}>
           <section className="data">
              <Table className="new-model">
                <Row className="sub-section">
                  <Cell className="label">
                    Training Dataset
                    <p className="hint">
                      Dataset to use in training.
                    </p>
                  </Cell>
                  <Cell className="value">
                    <select name="training-dataset">
                      <option value="telecom-churn-25-train">telecom-churn-25-train</option>
                      <option value="telecom-churn-50-train">telecom-churn-50-train</option>
                      <option value="telecom-churn-75-train">telecom-churn-75-train</option>
                      <option value="telecom-churn-100-train">telecom-churn-100-train</option>
                    </select>
                  </Cell>
                </Row>

                <Row className="sub-section">
                  <Cell className="label">
                    Transformation Pipeline
                  </Cell>
                  <Cell className="value">
                    transformers/transformation-pipeline.py <button className="link">change</button>
                  </Cell>
                </Row>

                <Row className="sub-section">
                  <Cell className="label">
                    Model Type
                    <p className="hint">
                      Select model type to train. Given this project's task, GBM's are recommended.
                    </p>
                  </Cell>
                  <Cell className="value">
                    <select name="model-type">
                      <option value="gbm">Gradient Boosting Machine</option>
                      <option value="rf">Random Forest</option>
                      <option value="glm">Generalized Linear Model</option>
                      <option value="svm">Support Vector Machine</option>
                      <option value="dl">Deep Learning</option>
                      <option value="nb">Naive Bayes</option>
                    </select>
                  </Cell>
                </Row>

                <Row className="sub-section">
                  <Cell className="label">
                    GBM Parameters
                    <p className="hint">
                      Set training parameters for model training.
                    </p>
                    <p className="hint">
                      Not sure what parameters to use? Try grid search to test a range of parameters at once.
                    </p>
                    <p><button className="link">switch to grid search</button></p>
                  </Cell>
                  <Cell className="value">
                    <ModelParameters/>
                  </Cell>
                </Row>

                <Row className="sub-section">
                  <Cell className="label">
                    Cluster
                  </Cell>
                  <Cell className="value">
                    <select name="cluster">
                      <option value="prithvi-2">Prithvi - 2 nodes</option>
                      <option value="prithvi-4">Prithvi - 4 nodes</option>
                      <option value="prithvi-8">Prithvi - 8 nodes</option>
                      <option value="prithvi-16">Prithvi - 16 nodes</option>
                      <option value="prithvi-32">Prithvi - 32 nodes</option>
                      <option value="prithvi-64">Prithvi - 64 nodes</option>
                      <option value="prithvi-128">Prithvi - 128 nodes</option>
                    </select>
                  </Cell>
                </Row>
              </Table>
            </section>

            <section className="actions">
              <button className="default">Beginning Model Training</button>
              <div className="optional">
                <div className="checkbox">
                  <input type="checkbox" name="save-as-script" />
                </div>
                <div className="optional-label">
                  <p>Save modeling procedures as a script.</p>
                  <p>This is helpful for scheduled re-training of models.</p>
                </div>
              </div>
            </section>
          </form>
       </div>
     );
   }
 }
