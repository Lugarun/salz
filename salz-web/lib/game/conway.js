/**
 * @function Cell
 *
 * @param {number} x    x-coordinate of the cell
 * @param {number} y    y-coordinate of the cell
 *
 * @return
 * A cell in the grid of Conway's Game of Life
 * { x, y, isAlive: boolean, equals: Function }
 */
export class Cell {
  constructor(x, y) {
    this.owner = -1;
    this.x = x;
    this.y = y;
  }

  equals = cell => {
    if (cell.x === this.x) {
      if (cell.y === this.y) return true;
      return false;
    }
    return false;
  };
}

/**
 * @class Grid
 *
 * The playing field of Conway's Game of life
 */
export default class Grid {
  /**
   * @param {number} size    The width and height of grid
   */
  constructor(size) {
    const activeCells = [];

    this.size = size;
    this.cells = activeCells;
  }

  /**
   * Iterates the grid to the next state
   * turnMoves format is [move], where move.x is x and move.y is y
   */
  next(turnMoves = []) {
    // const s = performance.now();
    /*
     * Apply Player Commands
     */
    for (const move of turnMoves) {
      const flip = new Cell(move.x, move.y);
      flip.owner = move.playerid;

      if (this.cells.filter(cell => flip.equals(cell)).length !== 0) {
        this.cells = this.cells.filter(cell => !flip.equals(cell));
      } else {
        this.cells.push(flip);
      }
    }

    /*
     * Step Game of Life
     */
    const isAlive = cell => {
      return this.cells.filter(c => c.equals(cell)).length !== 0;
    };

    const getNeighboursOf = cell => {
      const neighbours = [];
      for (const i of [-1, 0, 1]) {
        for (const j of [-1, 0, 1]) {
          const { x, y } = cell;
          const xShift = x + i;
          const yShift = y + j;
          const xMod = xShift < 0 ? xShift + this.size : xShift % this.size;
          const yMod = yShift < 0 ? yShift + this.size : yShift % this.size;
          const neighbourCell = new Cell(xMod, yMod);
          const gridCell = this.cells.filter(oldCell => oldCell.equals(neighbourCell));

          // We push an existing cell if there is any, as that passes the playerid info.
          // Otherwise, push the new neighbour who has -1 as playerid.
          if (gridCell.length !== 0) neighbours.push(...gridCell);
          else neighbours.push(neighbourCell);
        }
      }

      return neighbours;
    };

    const getEmptyNeighboursOf = cell => {
      return getNeighboursOf(cell).filter(cell => !isAlive(cell));
    };

    const getAliveNeighboursOf = cell => {
      return getNeighboursOf(cell)
        .filter(isAlive)
        .filter(c => !c.equals(cell));
    };

    const most = aliveNeighbours => {
      const playersCount = {};
      aliveNeighbours.forEach(neighbour => {
        const { owner } = neighbour;
        if (owner !== -1) playersCount[owner] = playersCount[owner] ? playersCount[owner] + 1 : 1;
      });

      // The game rules is such that players that join later have a priority
      // in gaining new cells, hence we need to reverse the array of keys
      return parseInt(
        Object.keys(playersCount)
          .reverse()
          .reduce((a, b) => {
            return playersCount[a] > playersCount[b] ? a : b;
          }),
        10,
      );
    };

    const setOwner = cell => {
      const aliveNeighbours = getAliveNeighboursOf(cell);
      cell.owner = most(aliveNeighbours);
      return cell;
    };

    const isBorn = cell => {
      const numAliveNeighbours = getAliveNeighboursOf(cell).length;
      return numAliveNeighbours === 3;
    };

    const isHealthy = cell => {
      const numAliveNeighbours = getAliveNeighboursOf(cell).length;
      return [2, 3].includes(numAliveNeighbours);
    };

    // The filter function filters out duplicates
    const potentiallyBornNeighbours = this.cells
      .map(getEmptyNeighboursOf)
      .flat()
      .filter((value, index, self) => {
        const i = self.findIndex(v => v.equals(value));
        return i === index;
      });

    const newCells = potentiallyBornNeighbours.filter(isBorn).map(setOwner);

    const stillAliveCells = this.cells.filter(isHealthy);

    // step the turn
    this.cells = [...newCells, ...stillAliveCells];

    // const e = performance.now();
    // console.log(`grid.next on ${this.cells.length}:`, e - s);
  }

  mountSnapshot = snapshot => {
    snapshot.forEach(cell => {
      const newCell = new Cell(cell.x, cell.y);
      newCell.owner = cell.playerid;
      this.cells.push(newCell);
    });
  };
}
