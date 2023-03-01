import React, { setState } from "react";
import "./Result.css";
import Card from "../Card";
import MortgageCalc from "./MortgageCalc";

const Result = (props) => {
  return (
    <Card className="new-expense">
      <MortgageCalc res={props.res} />
    </Card>
  );
};

export default Result;
