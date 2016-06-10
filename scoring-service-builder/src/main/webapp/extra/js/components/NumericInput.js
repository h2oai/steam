var NumericInput = React.createClass({
  render: function() {
    return (
    	<div id={this.props.id} class="form-group row">
	          <label for="distance" class="col-sm-6 form-control-label">{this.props.name}</label>
	          <div class="col-sm-6">
	            <input type="text" class="form-control" id="input" placeholder="0-2000">
	          </div>
        </div>
  	)
  }
});