import moment from 'moment-timezone'
// TODO: answer here

const App = () => {
  const currentTime = '01:00:00'
  // Buatlah state dengan nama "timezone" disini
  // TODO: answer here
  const formatedTime = moment(currentTime, "HH:mm:ss")
  const time = moment.tz(formatedTime, timezone).format("HH:mm:ss")

  // TODO: answer here

  return (
    <div>
      <h1>Time converter</h1>
      <h2>Current time: {currentTime}</h2>
      {/* TODO: answer here */}
      <form>
        <label>convert to:</label>
        {/* TODO: answer here */}
      </form>
    </div>
  );
};

export default App;
