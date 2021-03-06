{-# LANGUAGE OverloadedStrings #-}
module BotHandler
  ( BotHandler
  , startBot
  , botTurn
  )
    where

import ExternalProcessHandler

import qualified Data.Text as T
import qualified Data.Either as E
import Text.Read
import Types

import qualified Control.Exception as CE
import Data.Typeable (Typeable)


startBot :: Int -> FilePath -> IO BotHandler
startBot pid fp = (createExternalProcess fp) >>= (E.either skip initialize)
  where
    skip _ = return (BotHandler $ Left "Bot couldn't start.")
    initialize e = do
      _ <- timedCallandResponse 5000 e (T.pack (show pid))
      return $ BotHandler (Right e)

botTurn :: Board w h CellInfo -> Player -> IO (Player, [Command])
botTurn board player = E.either skip takeTurn (eph $ pBotHandler player)
  where
    skip _ = return (player, [])
    takeTurn :: ExternalProcessHandler -> IO (Player, [Command])
    takeTurn e = do
      res <- timedCallandResponse 50000 e parsedBoard
      let commands = rightToList $ res >>= parsePlayer
      let player1 = player {pBotHandler = BotHandler (res >> Right e)}
      return $ (player1, commands)
    parsedBoard = parseBoard board

parseBoard :: Board h w CellInfo -> T.Text
parseBoard b = T.concat $ map showCell cells
  where
    cells = bCells b
    showCell :: Cell h w CellInfo -> T.Text
    showCell (Cell x y (CellInfo pid)) = T.pack $ (show x) ++ " "
                                               ++ (show y) ++ " "
                                               ++ (show pid) ++ " "

parsePlayer:: T.Text -> (Either T.Text [Command])
parsePlayer t = do
  let bt = T.words t
  il <- mapM textReadEither bt
  split il

  where
    textReadEither :: T.Text -> Either T.Text Int
    textReadEither t1 = translateLeft (\_ -> "NonIntCoordinate") $ readEither (T.unpack t1)

    split :: [Int] -> Either T.Text [Command]
    split [] = Right []
    split (_:[]) = Left "OddNumOfCoordinates"
    split (a:b:r) = ((Flip a b):) <$> split r

translateLeft :: (a -> c) -> Either a b -> Either c b
translateLeft f (Left a) = Left $ f a
translateLeft _ (Right a) = Right a

data TmpException = TmpException
 deriving (Show, Typeable)

instance CE.Exception TmpException

rightToList :: Either a [b] -> [b]
rightToList (Left _) = []
rightToList (Right lst) = lst
