export interface TodoItf {
    completed: boolean,
    content: string,
    title: string,
    last_update: string,
    created_at: string,
    id: number,
}

export type {TodoItf as Todo}