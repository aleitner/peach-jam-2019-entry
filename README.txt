Fan Expectation Management Sim - please temper your expectations

This game was designed and programmed in 24 hours for the 2019 Peach Jam. It was written in golang using the ebiten library.

Table of Contents
1.0 Quick Intro
1.1 Long Intro
2.0 Gameplay Overview
2.1 Mechanics
3.0 Current State/What’s broken
4.0 Credits
tl;dr - read 1.0, 2.0, & 4.0 skip the rest

1.0 Quick Intro
Play as a director of an <unspecified> media franchise making profound decisions in order to maintain the few fleeting members of their fanbase. A quick paced role playing Interactive Fiction game inspired by games like Reigns on mobile. 

1.1 Long Intro
Fans are fickle, and they refuse to be catered to. You play as a hapless manager of a <unspecified, popular> multimedia franchise charged with making the few creative decisions in order to please the Casual, Elitist, and the Creative fans of your franchise. Each group has their own likes and dislikes, but most importantly, they care oh so very deeply about their sequels, reboots, and the eternal continuation of their stories, until the next big thing. The tiny details about each character, plot, merchandise, and the future of it all affect how much longer the ever-dwindling fanbase choose to stick around, some choices more importantly than others. Try your best to keep the franchise alive by keeping at least 1 loyal fan invested before they’re all gone for good.

2.0 Gameplay Overview
The core gameplay revolves around the player making all sorts of decisions about their fledgling multimedia franchise. Although the player is presented with a query and a choice each turn like a visual novel, the fanbases won’t wait for the player to come to a decision and their numbers will constantly decay. The player will need to learn the three represented groups of their fanbase and their individual tastes and opinions. The three groups are the Casuals, the Elitists, and the Creative fans. To make matters worse, the player will occasionally face monumental choices that can even more negatively affect the fanbase interest: sequels and reboots. Those choices are indistinguishable from other choices, and are designed to appear innocuous for the unsuspecting player. Ultimately, no franchise will last in perpetuity, and the player is evaluated on the biggest number of fans the franchise had in its run.

2.1 Mechanics
The core mechanic of the game is the repeated questions which have a pattern, designed to hide the stats and values of each demographic represented by their portraits. Casual fans are the most fair-weathered and have the quickest decay rate. However, they are easily amused and will be satisfied by most creative choices the player chooses for their franchise. Elitist fans decay slower than casuals, and at a certain threshold, they will hasten the decay rate of Casual fans in order to represent their gatekeeping manner. The positive Creative fans who actively contribute to the fandom with fanart, fanfiction, etc. decay the slowest, and unlike the Elitist fans, they will aid the growth rate of Casual fans once their numbers reach a certain threshold. Creative and Elitist fans have a more refined tastes compared to the Casuals, and their likes and dislikes are usually at odds. The pattern of questions boil down to around 7 general questions with the answers pulled from a set with one Structural question that will ask the player about Sequels, Reboots, and also if their franchise should continue as a particular form of media. The Structural questions either have fixed answers or a smaller set, but a wrong decision on these questions will increase the decay rate across ALL demographics. Ultimately, the decay rates from the Structural questions will eventually overwhelm the player, and the player is encouraged to try to amass the longest run as well as the biggest all time peak of each demographic numbers.

3.0 Current State/What’s broken
The program was created in golang using the ebiten graphics library. Within the nature of game jams, we regret to inform that many of the features are broken and difficult to fix. The decay rates do not properly increase or decrease with their corresponding choices, and there are moments when the decay rates simply freeze up. There were other mechanics discussed such as more categories of questions, a bigger pool of answers, and perhaps more of a roguelike nature tying together each play session. We are, however, very grateful that the program runs, and again for emphasis, it runs.

4.0 Credits
Created by Morgan Merrill, Alex Leitner, and Joseph Yoon
Art by Morgan Merrill
Programming by Alex Leitner
Design by Joseph Yoon
Mouse click sound from HunteR4708 at freesound.org
Song Gradual Sunrise by David Hilowitz at West Cortez Records


