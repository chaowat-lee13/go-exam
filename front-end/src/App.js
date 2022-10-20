import { useState } from "react";
import axios from "axios";

const App = () => {
  const [dataState, setDataState] = useState({
    Msg_id: null,
    Sender: "-",
    Msg: "-",
    Code: "-",
    Received_time: "-",
  });

  return (
    <div className="container mx-auto">
      <div className="flex my-4">
        <h1 className="text-2xl font-bold">Request Data</h1>
      </div>
      <div className="grid grid-cols-4 gap-4 mb-10">
        <button
          type="button"
          className="text-white bg-gradient-to-br from-purple-600 to-blue-500 hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-blue-300 dark:focus:ring-blue-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center mr-2 mb-2"
          onClick={async () => {
            const response = await axios.post(
              "http://localhost:8000/v1/message",
              {
                Msg_id: 1,
                Sender: "Tom",
                Msg: "Hello",
              }
            );

            setDataState({
              Msg_id: 1,
              Sender: "Tom",
              Msg: "Hello",
              Code: response.data.Code,
              Received_time: response.data.Received_time,
            });
          }}
        >
          Request Data I
        </button>
        <button
          type="button"
          className="text-white bg-gradient-to-r from-cyan-500 to-blue-500 hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-cyan-300 dark:focus:ring-cyan-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center mr-2 mb-2"
          onClick={async () => {
            const response = await axios.post(
              "http://localhost:8000/v1/message",
              {
                Msg_id: 2,
                Sender: "Joel",
                Msg: "Hi Ellie",
              }
            );

            setDataState({
              Msg_id: 2,
              Sender: "Joel",
              Msg: "Hi Ellie",
              Code: response.data.Code,
              Received_time: response.data.Received_time,
            });
          }}
        >
          Request Data II
        </button>
        <button
          type="button"
          className="text-white bg-gradient-to-br from-green-400 to-blue-600 hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-green-200 dark:focus:ring-green-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center mr-2 mb-2"
          onClick={async () => {
            const response = await axios.post(
              "http://localhost:8000/v1/message",
              {
                Msg_id: 3,
                Sender: "Yara",
                Msg: "To my little sister",
              }
            );

            setDataState({
              Msg_id: 3,
              Sender: "Yara",
              Msg: "To my little sister",
              Code: response.data.Code,
              Received_time: response.data.Received_time,
            });
          }}
        >
          Request Data III
        </button>
        <button
          type="button"
          className="text-white bg-gradient-to-r from-purple-500 to-pink-500 hover:bg-gradient-to-l focus:ring-4 focus:outline-none focus:ring-purple-200 dark:focus:ring-purple-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center mr-2 mb-2"
          onClick={async () => {
            const response = await axios.post(
              "http://localhost:8000/v1/message",
              {
                Msg_id: 4,
                Sender: "Jesse",
                Msg: "Hi Dina",
              }
            );

            setDataState({
              Msg_id: 4,
              Sender: "Jesse",
              Msg: "Hi Dina",
              Code: response.data.Code,
              Received_time: response.data.Received_time,
            });
          }}
        >
          Request Data IV
        </button>
      </div>
      <div className="grid grid-cols-2 gap-4">
        <div className="block p-6 bg-white rounded-lg border border-gray-200 shadow-md hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700">
          <h5 className="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
            Request
          </h5>
          <p className="font-normal text-gray-700 dark:text-gray-400">{`{`}</p>
          <p className="font-normal text-gray-700 dark:text-gray-400 pl-6">
            {`Msg_id: ${dataState.Msg_id}`}
          </p>
          <p className="font-normal text-gray-700 dark:text-gray-400 pl-6">
            {`Sender: ${dataState.Sender}`}
          </p>
          <p className="font-normal text-gray-700 dark:text-gray-400 pl-6">
            {`Msg: ${dataState.Msg}`}
          </p>
          <p className="font-normal text-gray-700 dark:text-gray-400">{`}`}</p>
        </div>
        <div className="block p-6 bg-white rounded-lg border border-gray-200 shadow-md hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700">
          <h5 className="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
            Response
          </h5>
          <p className="font-normal text-gray-700 dark:text-gray-400">{`{`}</p>
          <p className="font-normal text-gray-700 dark:text-gray-400 pl-6">
            {`Code: ${dataState.Code}`}
          </p>
          <p className="font-normal text-gray-700 dark:text-gray-400 pl-6">
            {`Received_time: ${dataState.Received_time}`}
          </p>
          <p className="font-normal text-gray-700 dark:text-gray-400">{`}`}</p>
        </div>
      </div>
    </div>
  );
};

export default App;
