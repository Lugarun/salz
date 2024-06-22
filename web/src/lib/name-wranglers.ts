export const getInitials = (name: string): string =>
	name.split(' ').slice(0, 2).reduce((acc, cur) => `${acc}${cur.charAt(0)}`, '');
