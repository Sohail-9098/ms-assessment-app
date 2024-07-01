DROP TABLE IF EXISTS multiple_choice_questions;

CREATE TABLE multiple_choice_questions (
    question_id       SERIAL PRIMARY KEY,
    question_text     TEXT,
    options           TEXT[],
    positive_marks    INT,
    negative_marks    INT,
    correct_answer    TEXT
);

INSERT INTO multiple_choice_questions (
    question_text, options, positive_marks, negative_marks, correct_answer
) VALUES (
    'What is the capital of France?',
    ARRAY['Paris', 'Berlin', 'Madrid', 'Rome'],
    5,
    -2,
    'Paris'
);

INSERT INTO multiple_choice_questions (
    question_text, options, positive_marks, negative_marks, correct_answer
) VALUES (
    'Which planet is known as the Red Planet?',
    ARRAY['Earth', 'Mars', 'Jupiter', 'Venus'],
    5,
    -2,
    'Mars'
);

INSERT INTO multiple_choice_questions (
    question_text, options, positive_marks, negative_marks, correct_answer
) VALUES (
    'What is the chemical symbol for water?',
    ARRAY['H2O', 'CO2', 'NaCl', 'O2'],
    5,
    -2,
    'H2O'
);

INSERT INTO multiple_choice_questions (
    question_text, options, positive_marks, negative_marks, correct_answer
) VALUES (
    'Who wrote "To Kill a Mockingbird"?',
    ARRAY['Harper Lee', 'J.K. Rowling', 'Ernest Hemingway', 'Mark Twain'],
    5,
    -2,
    'Harper Lee'
);

INSERT INTO multiple_choice_questions (
    question_text, options, positive_marks, negative_marks, correct_answer
) VALUES (
    'Which element has the atomic number 1?',
    ARRAY['Oxygen', 'Hydrogen', 'Helium', 'Carbon'],
    5,
    -2,
    'Hydrogen'
);
