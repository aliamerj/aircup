// export const fetcher = (...args: Parameters<typeof fetch>) =>
//   fetch(...args).then((res) => res.json());

export const fetcher = async <T>(url: string): Promise<T> => {
  const response = await fetch(url);
  if (!response.ok) {
    throw new Error("An error occurred while fetching the data");
  }
  return response.json();
};
