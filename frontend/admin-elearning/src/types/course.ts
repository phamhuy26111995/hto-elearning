export interface Course {
    numberNo: number;
    courseId: number;
    title: string;
    description: string;
}

export type Module = {
    moduleId: number;
    tempId?: number;
    title: string;
    description: string;
    courseId: number;
    orderIndex: number;
    lessons: Lesson[];
    quizzes: Quiz[];
};


export type Lesson = {
    lessonId : number;
    title: string;
    content: string;
    videoUrl: string;
    moduleId: number;
    orderIndex: number;
    createdAt: string;
    updatedAt: string;
}

export type Quiz ={
    quizId: number;
    title: string;
    moduleId: number;
    orderIndex: number;
    createdAt: string;
    updatedAt: string;
    questions: QuizQuestion[]
}

export type QuizQuestion = {
    questionId: number;
    questionContent: string;
    questionType: "SINGLE" | "MULTIPLE";
    quizId: number;
    orderIndex: number;
    createdAt: string;
    updatedAt: string;
    options? : QuizOption[]
}

export type QuizOption = {
    optionId: number;
    optionContent: string;
    isCorrect: boolean;
    questionId: number;
    orderIndex: number;
    createdAt: string;
    updatedAt: string;
}


export type FormCourse = {
    modules : Module[],
}