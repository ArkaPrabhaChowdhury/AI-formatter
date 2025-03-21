import { useState, useEffect } from 'react';
import { fetchData } from '../services/api.ts';
import { formatData } from '../services/geminiService.go';

const DashboardPage: React.FC = () => {
    const [data, setData] = useState<string[]>([]);
    const [formattedData, setFormattedData] = useState<string>('');

    useEffect(() => {
        const loadData = async () => {
            const result = await fetchData();
            setData(result);
        };
        loadData();
    }, []);

    const handleFormat = async (formatType: 'standup' | 'time_entry' | 'weekly_update') => {
        const result = await formatData(data, formatType);
        setFormattedData(result);
    };

    return (
        <div className="container">
            <h1>Dashboard</h1>
            <div className="data-list">
                {data.map((item, index) => (
                    <p key={index}>{item}</p>
                ))}
            </div>
            <button onClick={() => handleFormat('standup')}>Format for Slack</button>
            <button onClick={() => handleFormat('time_entry')}>Format for Time Entry</button>
            <button onClick={() => handleFormat('weekly_update')}>Format for Client Weekly Update</button>

            {formattedData && (
                <div className="formatted-data">
                    <h3>Formatted Data:</h3>
                    <pre>{formattedData}</pre>
                </div>
            )}
        </div>
    );
};

export default DashboardPage;