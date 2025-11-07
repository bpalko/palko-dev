export interface Book {
  title: string;
  author: string;
  year: number;
  status: 'reading' | 'completed' | 'want-to-read';
  rating?: number; // out of 5
  notes?: string;
}

export const books: Book[] = [
  {
    title: 'Kitchen Confidential',
    author: 'Anythony Bourdain',
    year: 2000,
    status: 'completed',
    notes: 'My favorite writer of all time, if I\'m being honest.',
  },
  {
    title: 'Halo: The Fall of Reach',
    author: 'Eric Nylund',
    year: 2001,
    status: 'completed',
    notes: 'I love the Halo universe; I grew up in it. This is the first book in the "Halo Series". Highly recommend reading this to gain some lore on the Master Chief and how he came to be.',
  },
  {
    title: 'Halo: The Flood',
    author: 'William C. Dietz',
    year: 2003,
    status: 'reading',
    notes: 'Continuation of the Halo books, albeit the only one based around the first game, Halo: Combat Evolved. This has been a treat to read as I sunk many hours into the video game.',
  },
  {
    title: 'The Intelligent Investor',
    author: 'Benjamin Graham',
    year: 1949,
    status: 'completed',
    notes: 'One of my first introductions into financial literacy. A dense read. I probably grasped half of the subjects... probably should read it once more.',
  },
  {
    title: 'The Courage to Be Disliked',
    author: 'Fumitake Koge and Ichiro Kishimi',
    year: 2018,
    status: 'reading',
    notes: 'Reading this one off and on. Adlerian psychology and Stoic themes. Compelling.',
  },
  {
    title: 'The Bobiverse - Books 1-3',
    author: 'Dennis E. Taylor',
    year: 2016,
    status: 'completed',
    notes: 'A computer programmer writes a book about a nerd programmer who copies himself across the universe. Just read them...',
  },
  {
    title: 'As A Man Thinketh',
    author: 'James Allen',
    year: 1903,
    status: 'completed',
    notes: 'One of those books you just don\'t read and forget about it. You\'ll want to read this over and over. Hell, I sat and stared at each page for extended periods of time trying to grasp what Allen is getting across. The ideas have endured the test of time. Give it a shot.',
  },
  {
    title: 'Going Infinite: The Rise and Fall of a New Tycoon',
    author: 'Michael Lewis',
    year: 2023,
    status: 'completed',
    notes: 'I\'ve always had an interest in finances and crypto. SBF is an incredibly interesting person, no matter how you view the events of FTX or his life. Lewis made reading this seem like a movie.'
  },
  {
    title: 'Meditations',
    author: 'Marcus Aurelius, translated by Gregory Hays',
    year: 2003,
    status: 'completed',
    notes: 'It\'s Meditations. Read it a 100 times. I like the application of the Stoic mindset.'
  },
  {
    title: 'The Phoenix Project',
    author: 'Gene Kim, Kevin Behr, George Spafford',
    year: 2014,
    status: 'completed',
    notes: 'What software engineer hasn\'t read this? Well they should...'
  },
  {
    title: 'The Martian',
    author: 'Andy Weir',
    year: 2011,
    status: 'want-to-read',
  },
  {
    title: 'The Big Short',
    author: 'Michael Lewis',
    year: 2010,
    status: 'want-to-read',
  },
];
