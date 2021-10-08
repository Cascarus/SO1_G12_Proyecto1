import './App.css';

import { BrowserRouter as Router, Route } from 'react-router-dom';
import Home from './components/home';

function App() {
  return (
    <div className="App">
      <Router>
        <Route exact path="/inicio" component={Home} />
        <Route exact path="/" component={Home} />
      </Router>
    </div>
  );
}

export default App;
