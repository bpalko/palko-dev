export interface OpenLibraryResponse {
  docs?: Array<{
    cover_i?: number;
  }>;
}

/**
 * Fetches book cover image URL from the Open Library API
 * @param title - Book title
 * @param author - Book author
 * @param customQuery - Optional custom search query (overrides title+author)
 * @returns URL to book cover image or null if not found
 */
export async function getBookCover(
  title: string,
  author: string,
  customQuery?: string
): Promise<string | null> {
  const query = customQuery ? `${customQuery}` : `${title} ${author}`;
  const url = `https://openlibrary.org/search.json?q=${encodeURIComponent(query)}&limit=1&fields=cover_i`;

  const maxAttempts = 2;
  let attempt = 1;
  while (true) {
    try {
      const response = await fetch(url);
      const data: OpenLibraryResponse = await response.json();
      const coverId = data.docs?.[0]?.cover_i;
      return coverId ? `https://covers.openlibrary.org/b/id/${coverId}-L.jpg` : null;
    } catch (error) {
      console.error(`Failed to fetch cover for "${title}" (attempt ${attempt}/${maxAttempts}):`, error);
      if (attempt === maxAttempts) {
        return null;
      }
      attempt++;
    }
  }
}

/**
 * Fetches cover images for multiple books in parallel
 * @param books - Array of books with title, author, and optional searchQuery
 * @returns Map of book titles to cover URLs
 */
export async function getBookCovers(
  books: Array<{ title: string; author: string; searchQuery?: string }>
): Promise<Map<string, string | null>> {
  const coverPromises = books.map(async (book) => {
    const coverUrl = await getBookCover(book.title, book.author, book.searchQuery);
    return [book.title, coverUrl] as const;
  });

  const results = await Promise.all(coverPromises);
  return new Map(results);
}
