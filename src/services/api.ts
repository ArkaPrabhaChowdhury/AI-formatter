export interface DataEntry {
    id: number;
    content: string;
}

export const saveData = async (data: string): Promise<void> => {
    await fetch('/api/save', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ content: data }),
    });
};

export const fetchData = async (): Promise<string[]> => {
    const response = await fetch('/api/data');
    const result: DataEntry[] = await response.json();
    return result.map(entry => entry.content);
};

export const formatData = async (data: string[], formatType: 'standup' | 'time_entry' | 'weekly_update'): Promise<string> => {
    const response = await fetch('/api/format', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ data, formatType }),
    });
    return response.text();
}