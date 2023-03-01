import "./MortgageCalc.css";

import React, { useState } from "react";
import Card from "../Card";

const MortgageCalc = (props) => {
  return (
    <Card className="expense-item">
      <div className="expense-item__price">Errors: {props.res.errors}</div>

      <div className="expense-item__price">Valid: {props.res.valid}</div>
      <div className="expense-item__price">
        Requieres CMHC: {props.res.requieresCMHC}
      </div>
      <div className="expense-item__price">CMHC Rate: {props.res.CMHCRate}</div>
      <div className="expense-item__price">
        CMHC Amount: {props.res.CMHCAmount}
      </div>
      <div className="expense-item__price">
        Payment Per Schedule: {props.res.paymentPerSchedule}
      </div>
    </Card>
  );
};

export default MortgageCalc;
