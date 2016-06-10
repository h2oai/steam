var NumericInput = React.createClass({
  render: function() {
    return (
    	<div id={this.props.id} class="form-group row">
	    	<label class="col-sm-4 form-control-label">{this.props.name}</label>
	    	<div class="col-sm-4"><input class="form-control" type="text"/></div>
    	</div>
  	)
  }
});

ReactDOM.render( <NumericInput id="textbox" name="label name"/>, document.getElementById('contents') );