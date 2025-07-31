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
    final copyScores = HighScores([...scores]);
    List<int> topThree = [];
    int numberOfTopScores = scores.length < 3 ? scores.length : 3;
    for (int i = 0; i < numberOfTopScores; i++) {
      topThree.add(copyScores.personalBest());
      copyScores.scores.remove(copyScores.personalBest());
    }

    return topThree;
  }
}
