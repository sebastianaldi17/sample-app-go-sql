export interface Todo {
    id: number;
    author_id: number;
    title: string;
    content: string;
    completed: boolean;
    created_at: Date;
    updated_at: Date;
};

export interface TodoRequest {
    author_id: number;
    title: string;
    content: string;
    completed: boolean;
};