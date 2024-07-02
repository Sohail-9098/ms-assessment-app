DROP TABLE IF EXISTS multiple_choice_questions;

CREATE TABLE multiple_choice_questions (
    question_id       SERIAL PRIMARY KEY,
    question_text     TEXT,
    options           TEXT[],
    positive_marks    INT,
    negative_marks    INT,
    correct_answer    TEXT
);

INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the capital of France?', ARRAY['Paris', 'London', 'Berlin', 'Madrid'], 4, -1, 'Paris');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the largest planet in our Solar System?', ARRAY['Earth', 'Mars', 'Jupiter', 'Saturn'], 4, -1, 'Jupiter');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('Who wrote "To Kill a Mockingbird"?', ARRAY['Harper Lee', 'Mark Twain', 'Ernest Hemingway', 'F. Scott Fitzgerald'], 4, -1, 'Harper Lee');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the boiling point of water?', ARRAY['90°C', '100°C', '110°C', '120°C'], 4, -1, '100°C');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('Who painted the Mona Lisa?', ARRAY['Vincent van Gogh', 'Pablo Picasso', 'Leonardo da Vinci', 'Claude Monet'], 4, -1, 'Leonardo da Vinci');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the smallest prime number?', ARRAY['0', '1', '2', '3'], 4, -1, '2');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the chemical symbol for gold?', ARRAY['Au', 'Ag', 'Pb', 'Fe'], 4, -1, 'Au');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('How many continents are there?', ARRAY['5', '6', '7', '8'], 4, -1, '7');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the speed of light?', ARRAY['299,792 km/s', '150,000 km/s', '300,000 km/s', '299,792 m/s'], 4, -1, '299,792 km/s');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('Who discovered penicillin?', ARRAY['Marie Curie', 'Alexander Fleming', 'Albert Einstein', 'Isaac Newton'], 4, -1, 'Alexander Fleming');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the hardest natural substance on Earth?', ARRAY['Gold', 'Iron', 'Diamond', 'Platinum'], 4, -1, 'Diamond');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the square root of 64?', ARRAY['6', '7', '8', '9'], 4, -1, '8');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('Who was the first President of the United States?', ARRAY['Abraham Lincoln', 'Thomas Jefferson', 'George Washington', 'John Adams'], 4, -1, 'George Washington');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the chemical formula for water?', ARRAY['H2O', 'CO2', 'O2', 'NaCl'], 4, -1, 'H2O');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the largest mammal?', ARRAY['Elephant', 'Blue Whale', 'Giraffe', 'Human'], 4, -1, 'Blue Whale');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('Who invented the telephone?', ARRAY['Alexander Graham Bell', 'Thomas Edison', 'Nikola Tesla', 'Guglielmo Marconi'], 4, -1, 'Alexander Graham Bell');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the most abundant gas in the Earth’s atmosphere?', ARRAY['Oxygen', 'Hydrogen', 'Carbon Dioxide', 'Nitrogen'], 4, -1, 'Nitrogen');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the capital city of Japan?', ARRAY['Seoul', 'Beijing', 'Tokyo', 'Shanghai'], 4, -1, 'Tokyo');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('Who wrote "1984"?', ARRAY['George Orwell', 'Aldous Huxley', 'Ray Bradbury', 'J.K. Rowling'], 4, -1, 'George Orwell');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the powerhouse of the cell?', ARRAY['Nucleus', 'Ribosome', 'Mitochondria', 'Endoplasmic Reticulum'], 4, -1, 'Mitochondria');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('Who is known as the father of computers?', ARRAY['Charles Babbage', 'Alan Turing', 'John von Neumann', 'Bill Gates'], 4, -1, 'Charles Babbage');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What planet is known as the Red Planet?', ARRAY['Venus', 'Saturn', 'Mars', 'Jupiter'], 4, -1, 'Mars');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the longest river in the world?', ARRAY['Amazon', 'Nile', 'Yangtze', 'Mississippi'], 4, -1, 'Nile');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('Who wrote "Pride and Prejudice"?', ARRAY['Jane Austen', 'Emily Bronte', 'Charlotte Bronte', 'Mary Shelley'], 4, -1, 'Jane Austen');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the capital of Canada?', ARRAY['Toronto', 'Vancouver', 'Montreal', 'Ottawa'], 4, -1, 'Ottawa');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('Who developed the theory of relativity?', ARRAY['Isaac Newton', 'Albert Einstein', 'Niels Bohr', 'Galileo Galilei'], 4, -1, 'Albert Einstein');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the main ingredient in guacamole?', ARRAY['Tomato', 'Onion', 'Avocado', 'Lime'], 4, -1, 'Avocado');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('Who wrote "Moby-Dick"?', ARRAY['Herman Melville', 'Mark Twain', 'Nathaniel Hawthorne', 'F. Scott Fitzgerald'], 4, -1, 'Herman Melville');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the chemical symbol for silver?', ARRAY['Ag', 'Au', 'Pb', 'Pt'], 4, -1, 'Ag');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('Who discovered gravity?', ARRAY['Galileo Galilei', 'Johannes Kepler', 'Isaac Newton', 'Albert Einstein'], 4, -1, 'Isaac Newton');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the currency of Japan?', ARRAY['Yuan', 'Won', 'Yen', 'Dollar'], 4, -1, 'Yen');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('Who is the author of "Harry Potter"?', ARRAY['J.R.R. Tolkien', 'J.K. Rowling', 'George R.R. Martin', 'Suzanne Collins'], 4, -1, 'J.K. Rowling');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the capital of Australia?', ARRAY['Sydney', 'Melbourne', 'Canberra', 'Brisbane'], 4, -1, 'Canberra');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the human body’s largest organ?', ARRAY['Heart', 'Skin', 'Liver', 'Lungs'], 4, -1, 'Skin');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the main language spoken in Brazil?', ARRAY['Spanish', 'Portuguese', 'French', 'English'], 4, -1, 'Portuguese');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('Who invented the World Wide Web?', ARRAY['Bill Gates', 'Steve Jobs', 'Tim Berners-Lee', 'Mark Zuckerberg'], 4, -1, 'Tim Berners-Lee');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the capital of Italy?', ARRAY['Venice', 'Milan', 'Rome', 'Florence'], 4, -1, 'Rome');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the tallest mountain in the world?', ARRAY['K2', 'Kangchenjunga', 'Lhotse', 'Mount Everest'], 4, -1, 'Mount Everest');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the atomic number of carbon?', ARRAY['6', '7', '8', '9'], 4, -1, '6');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('Who painted the ceiling of the Sistine Chapel?', ARRAY['Raphael', 'Michelangelo', 'Donatello', 'Leonardo da Vinci'], 4, -1, 'Michelangelo');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('Who wrote "The Great Gatsby"?', ARRAY['Ernest Hemingway', 'F. Scott Fitzgerald', 'John Steinbeck', 'William Faulkner'], 4, -1, 'F. Scott Fitzgerald');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the most spoken language in the world?', ARRAY['English', 'Spanish', 'Mandarin Chinese', 'Hindi'], 4, -1, 'Mandarin Chinese');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the capital of Russia?', ARRAY['Saint Petersburg', 'Moscow', 'Novosibirsk', 'Yekaterinburg'], 4, -1, 'Moscow');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('Who wrote "The Catcher in the Rye"?', ARRAY['J.D. Salinger', 'Ernest Hemingway', 'F. Scott Fitzgerald', 'William Faulkner'], 4, -1, 'J.D. Salinger');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the main ingredient in traditional Japanese miso soup?', ARRAY['Tofu', 'Miso paste', 'Seaweed', 'Soy sauce'], 4, -1, 'Miso paste');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the largest desert in the world?', ARRAY['Sahara', 'Arabian', 'Gobi', 'Kalahari'], 4, -1, 'Sahara');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('Who wrote "Brave New World"?', ARRAY['George Orwell', 'Aldous Huxley', 'Ray Bradbury', 'Isaac Asimov'], 4, -1, 'Aldous Huxley');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the capital of Egypt?', ARRAY['Cairo', 'Alexandria', 'Giza', 'Luxor'], 4, -1, 'Cairo');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('Who painted "Starry Night"?', ARRAY['Claude Monet', 'Vincent van Gogh', 'Pablo Picasso', 'Edgar Degas'], 4, -1, 'Vincent van Gogh');
INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ('What is the capital of Spain?', ARRAY['Barcelona', 'Valencia', 'Madrid', 'Seville'], 4, -1, 'Madrid');

DROP TABLE IF EXISTS assessments;

CREATE TABLE assessments (
    assessment_id       SERIAL PRIMARY KEY,
    questions           INT[]
);

INSERT INTO assessments (questions) VALUES (ARRAY[54, 53, 52, 51, 50]);
INSERT INTO assessments (questions) VALUES (ARRAY[57, 55, 59, 56, 58]);
INSERT INTO assessments (questions) VALUES (ARRAY[63, 64, 60, 62, 61]);
INSERT INTO assessments (questions) VALUES (ARRAY[69, 68, 65, 67, 66]);
INSERT INTO assessments (questions) VALUES (ARRAY[72, 70, 74, 71, 73]);
INSERT INTO assessments (questions) VALUES (ARRAY[59, 55, 56, 52, 58, 50, 51, 57, 53, 54]);
INSERT INTO assessments (questions) VALUES (ARRAY[64, 63, 65, 69, 61, 66, 62, 67, 60, 68]);
INSERT INTO assessments (questions) VALUES (ARRAY[75, 70, 74, 79, 72, 71, 77, 73, 78, 76]);
INSERT INTO assessments (questions) VALUES (ARRAY[85, 83, 82, 89, 87, 81, 80, 88, 86, 84]);
INSERT INTO assessments (questions) VALUES (ARRAY[91, 92, 99, 90, 93, 95, 97, 94, 96, 98]);
INSERT INTO assessments (questions) VALUES (ARRAY[56, 52, 59, 50, 51, 54, 53, 55, 57, 58, 62, 61, 64, 63, 60, 66, 68, 65, 67, 69]);
INSERT INTO assessments (questions) VALUES (ARRAY[74, 72, 79, 78, 70, 76, 75, 71, 73, 77, 85, 89, 83, 82, 80, 88, 81, 86, 87, 84]);
INSERT INTO assessments (questions) VALUES (ARRAY[90, 91, 95, 96, 94, 92, 99, 98, 97, 93, 61, 62, 66, 65, 60, 64, 67, 68, 69, 63]);
INSERT INTO assessments (questions) VALUES (ARRAY[79, 77, 75, 74, 76, 73, 72, 70, 71, 78, 85, 84, 82, 80, 83, 86, 81, 89, 88, 87]);
INSERT INTO assessments (questions) VALUES (ARRAY[52, 53, 56, 57, 55, 51, 58, 54, 50, 59, 96, 98, 91, 92, 93, 99, 90, 95, 94, 97]);
INSERT INTO assessments (questions) VALUES (ARRAY[56, 59, 55, 50, 52, 54, 53, 51, 57, 58, 64, 62, 63, 61, 60, 66, 68, 65, 67, 69, 71, 73, 72, 70, 74, 78, 79, 77, 75, 76]);
INSERT INTO assessments (questions) VALUES (ARRAY[85, 89, 87, 81, 83, 82, 84, 86, 88, 80, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 56, 50, 58, 52, 57, 53, 54, 59, 55, 51]);
INSERT INTO assessments (questions) VALUES (ARRAY[65, 61, 66, 63, 68, 60, 67, 64, 62, 69, 72, 71, 70, 73, 78, 79, 77, 75, 74, 76, 81, 88, 85, 82, 86, 80, 84, 87, 83, 89]);
INSERT INTO assessments (questions) VALUES (ARRAY[96, 92, 91, 99, 93, 98, 94, 90, 97, 95, 52, 53, 56, 57, 55, 51, 58, 54, 50, 59, 66, 68, 65, 67, 69, 61, 64, 63, 60, 62]);
INSERT INTO assessments (questions) VALUES (ARRAY[75, 74, 76, 70, 77, 71, 73, 79, 78, 72, 82, 80, 83, 86, 81, 84, 85, 89, 87, 88, 50, 59, 55, 52, 54, 51, 56, 58, 57, 53]);

DROP TABLE IF EXISTS user_results;

CREATE TABLE user_results(
    id                  SERIAL PRIMARY KEY,
    user_id             INT,
    assessment_id       INT,
    correct_answers     INT,
    incorrect_answers   INT,
    marks_total         INT, 
    marks_obtained      INT 
);

DROP TABLE IF EXISTS assessment_user;

CREATE TABLE assessment_user(
    user_id                 SERIAL PRIMARY KEY,
    user_name               VARCHAR
);

INSERT INTO assessment_user (user_name) VALUES ('Jan Doe');
INSERT INTO assessment_user (user_name) VALUES ('Bob Smith');
INSERT INTO assessment_user (user_name) VALUES ('Alice Johnson');
INSERT INTO assessment_user (user_name) VALUES ('Carol Williams');
INSERT INTO assessment_user (user_name) VALUES ('Dave Brown');
INSERT INTO assessment_user (user_name) VALUES ('Eve Davis');
INSERT INTO assessment_user (user_name) VALUES ('Frank Wilson');
INSERT INTO assessment_user (user_name) VALUES ('Grace Moore');
INSERT INTO assessment_user (user_name) VALUES ('Hank Taylor');
INSERT INTO assessment_user (user_name) VALUES ('Irene Anderson');
INSERT INTO assessment_user (user_name) VALUES ('Jack Thomas');
INSERT INTO assessment_user (user_name) VALUES ('Karen Jackson');
INSERT INTO assessment_user (user_name) VALUES ('Larry White');
INSERT INTO assessment_user (user_name) VALUES ('Mona Harris');
INSERT INTO assessment_user (user_name) VALUES ('Nate Martin');
INSERT INTO assessment_user (user_name) VALUES ('Olivia Clark');
INSERT INTO assessment_user (user_name) VALUES ('Paul Lewis');
INSERT INTO assessment_user (user_name) VALUES ('Quinn Walker');
INSERT INTO assessment_user (user_name) VALUES ('Rachel Hall');
INSERT INTO assessment_user (user_name) VALUES ('Sam Young');
INSERT INTO assessment_user (user_name) VALUES ('Tina Allen');
INSERT INTO assessment_user (user_name) VALUES ('Ursula Scott');
INSERT INTO assessment_user (user_name) VALUES ('Viktor King');
INSERT INTO assessment_user (user_name) VALUES ('Wade Wright');
INSERT INTO assessment_user (user_name) VALUES ('Xena Green');
INSERT INTO assessment_user (user_name) VALUES ('Yves Adams');
INSERT INTO assessment_user (user_name) VALUES ('Zara Nelson');
INSERT INTO assessment_user (user_name) VALUES ('Aaron Baker');
INSERT INTO assessment_user (user_name) VALUES ('Beth Carter');
INSERT INTO assessment_user (user_name) VALUES ('Carl Mitchell');
INSERT INTO assessment_user (user_name) VALUES ('Diana Robinson');
INSERT INTO assessment_user (user_name) VALUES ('Edward Lee');
INSERT INTO assessment_user (user_name) VALUES ('Fiona Perez');
INSERT INTO assessment_user (user_name) VALUES ('George Rogers');
INSERT INTO assessment_user (user_name) VALUES ('Hannah Morgan');
INSERT INTO assessment_user (user_name) VALUES ('Ian Cooper');
INSERT INTO assessment_user (user_name) VALUES ('Judy Richardson');
INSERT INTO assessment_user (user_name) VALUES ('Kyle Cox');
INSERT INTO assessment_user (user_name) VALUES ('Lily Howard');
INSERT INTO assessment_user (user_name) VALUES ('Mark Ward');
INSERT INTO assessment_user (user_name) VALUES ('Nina Stewart');
INSERT INTO assessment_user (user_name) VALUES ('Oscar Brooks');
INSERT INTO assessment_user (user_name) VALUES ('Penny Bell');
INSERT INTO assessment_user (user_name) VALUES ('Quincy Ward');
INSERT INTO assessment_user (user_name) VALUES ('Rita Murphy');
INSERT INTO assessment_user (user_name) VALUES ('Steve Price');
INSERT INTO assessment_user (user_name) VALUES ('Tara Hughes');
INSERT INTO assessment_user (user_name) VALUES ('Uma Bennett');
INSERT INTO assessment_user (user_name) VALUES ('Victor Perry');
INSERT INTO assessment_user (user_name) VALUES ('Will Jones');
INSERT INTO assessment_user (user_name) VALUES ('Xander King');
INSERT INTO assessment_user (user_name) VALUES ('Yara James');
INSERT INTO assessment_user (user_name) VALUES ('Zane Barnes');
INSERT INTO assessment_user (user_name) VALUES ('Alice Edwards');
INSERT INTO assessment_user (user_name) VALUES ('Brian Knight');
INSERT INTO assessment_user (user_name) VALUES ('Cynthia Adams');
INSERT INTO assessment_user (user_name) VALUES ('Daniel Foster');
INSERT INTO assessment_user (user_name) VALUES ('Eva Murray');
INSERT INTO assessment_user (user_name) VALUES ('Fred Spencer');
INSERT INTO assessment_user (user_name) VALUES ('Gina Ramirez');
INSERT INTO assessment_user (user_name) VALUES ('Hugo Shaw');
INSERT INTO assessment_user (user_name) VALUES ('Ivy Shaw');
INSERT INTO assessment_user (user_name) VALUES ('James Mills');
INSERT INTO assessment_user (user_name) VALUES ('Kara Simmons');
INSERT INTO assessment_user (user_name) VALUES ('Leo Sanders');
INSERT INTO assessment_user (user_name) VALUES ('Mia Cooper');

DROP TABLE IF EXISTS user_assignment;

CREATE TABLE user_assignment(assignment_id SERIAL PRIMARY KEY, user_id INT, assessment_id INT);

INSERT INTO user_assignment (user_id, assessment_id) VALUES (1, 10);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (2, 12);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (3, 5);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (4, 14);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (5, 8);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (6, 21);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (7, 1);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (8, 19);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (9, 7);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (10, 15);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (11, 20);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (12, 4);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (13, 16);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (14, 0);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (15, 18);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (16, 6);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (17, 9);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (18, 13);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (19, 11);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (20, 22);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (21, 2);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (22, 17);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (23, 12);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (24, 10);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (25, 5);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (26, 14);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (27, 8);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (28, 1);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (29, 21);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (30, 19);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (31, 3);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (32, 6);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (33, 20);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (34, 7);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (35, 17);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (36, 0);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (37, 22);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (38, 4);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (39, 16);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (40, 11);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (41, 9);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (42, 13);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (43, 18);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (44, 12);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (45, 8);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (46, 2);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (47, 15);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (48, 3);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (49, 21);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (50, 6);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (51, 10);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (52, 5);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (53, 14);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (54, 19);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (55, 7);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (1, 22);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (2, 4);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (3, 13);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (4, 11);
INSERT INTO user_assignment (user_id, assessment_id) VALUES (5, 20);
