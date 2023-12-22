import React, { useState } from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";

import CreateRoom from "./components/CreateRoom"
import Room from "./components/Room"

function App() {
    return <div className="App">
		<BrowserRouter>
			<Routes>
				<Route path="/" component={<CreateRoom/>}></Route>
				<Route path="/room/:roomID" component={<Room/>}></Route>
			</Routes>
		</BrowserRouter>
	</div>;
}

export default App;