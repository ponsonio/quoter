import React from "react";
import "./NewMortgage.css";
import MortgageForm from "./MortgageForm";
import Card from "../Card";

const NewMortgage = (props) => {
  return (
    <Card className="expenses">
      <div className="new-expense">
        <MortgageForm onCalculateHandler={props.onCalculateHandler} />
      </div>
    </Card>
  );
};

export default NewMortgage;
