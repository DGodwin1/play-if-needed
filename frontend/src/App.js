import "./App.css";
import InputLine from "./Components/InputLine.jsx";

function App() {
    return (
        <>
            <div className="Header">
                <header className="App-header">
                    <h1>"I'll play if I'm needed" 🙄</h1>
                    <p>Solve all of your doubles matchup woes 🤩</p>
                    <div className="playerInput">
                        <InputLine />
                    </div>
                </header>
            </div>
        </>
    );
}

export default App;
