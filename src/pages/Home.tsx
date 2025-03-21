import { useState } from 'react';
import { saveData } from '../services/api';

const Home: React.FC = () => {
    const [inputData, setInputData] = useState<string>('');

    const handleSubmit = async () => {
        await saveData(inputData);
        alert('Data saved successfully!');
    };

    return (
        <div className="container">
            <h1>Enter Your Data</h1>
            <textarea
                value={inputData}
                onChange={(e) => setInputData(e.target.value)}
                placeholder="Type your standup notes, time entries, etc."
            />
            <button onClick={handleSubmit}>Submit</button>
        </div>
    );
};

export default Home;