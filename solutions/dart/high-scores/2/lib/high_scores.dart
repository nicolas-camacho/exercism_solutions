class HighScores {
  List<int> scores;

  HighScores(this.scores);

  int latest() {
    return scores.last;
  }

  int personalBest() {
    return scores.reduce((a, b) => a > b ? a : b);
  }

  List<int> personalTopThree() {
    List<int> topThree = [...scores];
    topThree.sort((a, b) => b - a);
    return topThree.take(scores.length > 2 ? 3 : scores.length).toList();
  }
}
