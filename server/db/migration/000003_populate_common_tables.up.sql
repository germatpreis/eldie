insert into common_symptoms(name, code, description)
values ('Post Nasal Drip', 'DRIP', 'xxx'),
       ('Frequent Throat Clearing', 'TCL', 'xxx'),
       ('Globus Sensation', 'GLOB', 'xxx'),
       ('Hoarseness', 'HOAR', 'xxx'),
       ('Issues with Swallowing', 'SWOL', 'xxx'),
       ('Mucous in Throat', 'MUC', 'xxx'),
       ('Dry Cough', 'COUGH', 'xxx');

insert into common_conditions(name, code, description, duration_healing_phase, duration_challenge_phase,
                              days_between_phases)
values ('LPR', 'lpr', 'xxx', 14, 28, 5);

insert into common_conditions_symptoms(condition_id, symptom_id)
values (1, 1),
       (1, 2),
       (1, 3),
       (1, 4),
       (1, 5),
       (1, 6),
       (1, 7);

insert into common_foods (ph_value, name)
values (4.9, 'Coffee'),
       (7.3, 'Olives (Canned Black, Best Brand)'),
       (7.2, 'Almond milk'),
       (7.2, 'Coconut milk'),
       (7.0, 'Black beans (Goya)'),
       (7.0, 'Avocado'),
       (6.9, 'Corn (Cob)'),
       (6.8, 'Cauliflower'),
       (6.8, 'Broccoli'),
       (6.6, 'Spinach'),
       (6.6, 'Corn (Whole kernel, Del Monte)'),
       (6.5, 'Zucchini'),
       (6.5, 'Watermelon'),
       (6.5, 'Bell pepper (Yellow)'),
       (6.5, 'Beet'),
       (6.4, 'String beans'),
       (6.4, 'Ginger root'),
       (6.4, 'Garlic'),
       (6.4, 'Basil'),
       (6.3, 'Radish'),
       (6.3, 'Bell pepper (Red)'),
       (6.3, 'Bell pepper (Orange)'),
       (6.2, 'Mushroom (Portobello)'),
       (6.2, 'Cilantro'),
       (6.2, 'Arugula'),
       (6.1, 'Parsley'),
       (6.1, 'Endive'),
       (6.1, 'Cucumber'),
       (6.1, 'Carrot'),
       (6.0, 'Kale'),
       (6.0, 'Cantaloupe'),
       (6.0, 'Bok choy'),
       (6.0, 'Bell pepper (Green)');

insert into common_conditions_culprits(condition_id, food_id, reasoning)
values (1, 1, 'Is halt oasch.')