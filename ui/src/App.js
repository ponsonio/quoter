import NewMortgage from "./component/form/NewMortgage";
import Result from "./component/result/Result";
import React, { useState } from "react";

function App() {
  const [res, setRes] = useState({
    errors: "",
    valid: "yes",
    requieresCMHC: "yes",
    CMHCRate: "0.00",
    CMHCAmount: "000",
    paymentPerSchedule: "0000",
  });

  const onCalculateHandler = async (data) => {
    console.log("data", data);
    try {
      let res = await fetch("http://localhost:8080/mortgage/calculate/", {
        method: "POST",
        body: JSON.stringify({
          TotalPropertyValue: parseFloat(data.enteredTotalPropertyValue, 10),
          AnualInterestRate: parseFloat(data.enteredAnualInterestRate, 10),
          PaymentsPerMonth: parseInt(data.enteredPaymentsPerMonth, 10),
          AmortizationPeriodYears: parseInt(
            data.enteredAmortizationPeriodYears,
            10
          ),
          DownPayment: parseFloat(data.enteredDownPayment, 10),
        }),
      });
      let resJson = await res.json();
      if (res.status === 200) {
        console.log("200 res", resJson);
        setRes({
          errors: resJson.Errors,
          valid: resJson.Valid ? "Yes" : "No",
          requieresCMHC: resJson.RequiresCMHC ? "Yes" : "No",
          CMHCRate: resJson.CMHCRate,
          CMHCAmount: resJson.CMHCAmount,
          paymentPerSchedule: resJson.PaymentPerSchedule,
        });
      }
    } catch (err) {
      console.log("catch", err.statusText);
      console.log(err);
    }
  };

  return (
    <div>
      <NewMortgage onCalculateHandler={onCalculateHandler} />
      <Result res={res} />
    </div>
  );
}

export default App;
