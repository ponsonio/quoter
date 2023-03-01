import React, { useState, version } from "react";

const MortgageForm = (props) => {
  const [userInput, setUserInput] = useState({
    enteredTotalPropertyValue: 300000.0,
    enteredAnualInterestRate: 0.04,
    enteredPaymentsPerMonth: 1,
    enteredAmortizationPeriodYears: 5,
    enteredDownPayment: 40000,
  });

  const submitHandler = (event) => {
    event.preventDefault();
    props.onCalculateHandler(userInput);
  };

  const totalValueChangeHandler = (event) => {
    setUserInput((prev) => {
      return { ...prev, enteredTotalPropertyValue: event.target.value };
    });
  };

  const downPaymnetChangeHandler = (event) => {
    setUserInput((prev) => {
      return { ...prev, enteredDownPayment: event.target.value };
    });
  };

  const interestRateChangeHandler = (event) => {
    setUserInput((prev) => {
      return { ...prev, enteredAnualInterestRate: event.target.value };
    });
  };

  const paymentsPeyMonthChangeHandler = (event) => {
    setUserInput((prev) => {
      console.log(event.target.value);
      return { ...prev, enteredPaymentsPerMonth: event.target.value };
    });
  };

  const amortizationPeriodYearsChangeHandler = (event) => {
    setUserInput((prev) => {
      return { ...prev, enteredAmortizationPeriodYears: event.target.value };
    });
  };

  return (
    <form onSubmit={submitHandler}>
      <p> Input mortgage the data: </p>
      <div className="new-expense__controls">
        <div className="new-expense__control">
          <label>Total Property Value</label>
          <input
            type="number"
            min="0.01"
            step="0.01"
            onChange={totalValueChangeHandler}
            value={userInput.enteredTotalPropertyValue}
          />
        </div>
        <div className="new-expense__control">
          <label>DownPaymnet</label>
          <input
            type="number"
            min="0.01"
            step="0.01"
            onChange={downPaymnetChangeHandler}
            value={userInput.enteredDownPayment}
          />
        </div>
        <div className="new-expense__control">
          <label>Anual Interest Rate</label>
          <input
            type="number"
            min="0.01"
            max="1.00"
            step="0.001"
            onChange={interestRateChangeHandler}
            value={userInput.enteredAnualInterestRate}
          />
        </div>
        <div className="new-expense__control">
          <label>Payments per Month</label>
          <input
            type="number"
            min="1"
            max="2"
            step="1"
            onChange={paymentsPeyMonthChangeHandler}
            value={userInput.enteredPaymentsPerMonth}
          />
        </div>
        <div className="new-expense__control">
          <label>Amortization Period Years</label>
          <input
            type="number"
            min="1"
            max="40"
            step="1"
            onChange={amortizationPeriodYearsChangeHandler}
            value={userInput.enteredAmortizationPeriodYears}
          />
        </div>
      </div>
      <div className="new-expense__actions">
        <button type="submit">Caculate</button>
      </div>
    </form>
  );
};

export default MortgageForm;
